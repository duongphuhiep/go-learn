package middlewares

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
	ctx        context.Context
	url        *url.URL
	logger     *slog.Logger
}

var _ http.ResponseWriter = (*responseWriter)(nil)

const maxResponseBodySize int = 1024 * 2  // 2 KB
const maxResponseBodyTrunkSize int = 1024 // 1 KB

// const maxResponseBodySize int = 256
// const maxResponseBodyTrunkSize int = 64

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.logger.InfoContext(rw.ctx, fmt.Sprintf("HTTP Response status %d %s", code, rw.url),
		appendHeader([]any{
			"response.status", code,
		}, "response.", rw.Header())...,
	)
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	var trunkContent string
	if len(b) > maxResponseBodyTrunkSize {
		trunkContent = (string)(b[:maxResponseBodyTrunkSize]) + "... [truncated]"
	} else {
		trunkContent = string(b)
	}

	// Prepare header fields for structured logging
	rw.logger.InfoContext(rw.ctx, fmt.Sprintf("HTTP Response trunk %s", rw.url),
		appendHeader([]any{
			"response.status", rw.statusCode,
			"response.body", trunkContent,
		}, "response.", rw.Header())...,
	)

	// Only buffer a limited amount for SSE to avoid memory issues
	if rw.body.Len() < maxResponseBodySize {
		remainingSpace := maxResponseBodySize - rw.body.Len()
		if len(b) <= remainingSpace {
			rw.body.Write(b)
		} else {
			rw.body.Write(b[:remainingSpace])
			rw.body.WriteString("... [truncated]")
		}
	}

	return rw.ResponseWriter.Write(b)
}

// Flush implements http.Flusher for SSE support
func (rw *responseWriter) Flush() {
	if flusher, ok := rw.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}

func appendHeader(initAttrs []any, prefix string, header http.Header) []any {
	for k, v := range header {
		initAttrs = append(initAttrs, prefix+k, v)
	}
	return initAttrs
}

// loggingMiddleware creates a middleware that logs requests and responses with SSE support
func LoggingMiddleware(next http.Handler) http.Handler {
	tracer := otel.Tracer("http-middleware")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start a new span for this request
		ctx, span := tracer.Start(r.Context(), "http_request",
			oteltrace.WithAttributes(
				attribute.String("http.method", r.Method),
				attribute.String("http.url", r.URL.String()),
				attribute.String("http.scheme", r.URL.Scheme),
				attribute.String("http.host", r.Host),
				attribute.String("http.target", r.URL.Path),
			),
		)
		defer span.End()

		// Prepare request headers as logger attributes with "requestHeader." prefix
		logger := slog.With(appendHeader([]any{
			"request.method", r.Method,
			"request.url", r.URL.String(),
		}, "request.", r.Header)...)

		// Read and log request body
		var requestBody []byte
		if r.Body != nil {
			requestBody, _ = io.ReadAll(r.Body)
			// Restore the body for the next handler
			r.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// Log the request with correlation ID
		logger.InfoContext(ctx, fmt.Sprintf("HTTP Request %s  %s", r.Method, r.URL),
			"request.body", string(requestBody),
		)
		//slog.InfoContext(ctx, fmt.Sprintf("HTTP Request with defaultLogger %s  %s", r.Method, r.URL))

		// Add span attributes for request body
		if len(requestBody) > 0 {
			span.SetAttributes(attribute.String("http.request.body", string(requestBody)))
		}

		// Wrap the ResponseWriter to capture response data
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // default status
			body:           &bytes.Buffer{},
			ctx:            ctx,
			url:            r.URL,
			logger:         logger,
		}

		// Pass the context with span to the next handler
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(rw, r)

		// Add response attributes to span
		span.SetAttributes(
			attribute.Int("http.status_code", rw.statusCode),
			attribute.String("http.response.body", rw.body.String()),
		)

		// Log the response with correlation ID
		logger.InfoContext(ctx, fmt.Sprintf("HTTP response final %s", r.URL),
			appendHeader([]any{
				"response.status", rw.statusCode,
				"response.body", rw.body.String(),
			}, "response.", rw.Header())...,
		)
		// Set span status based on HTTP status code
		if rw.statusCode >= 400 {
			span.SetAttributes(attribute.Bool("error", true))
		}
	})
}
