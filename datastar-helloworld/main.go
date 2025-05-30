package main

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	otlogsdk "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	_ "embed"
	"net/http"

	m "github.com/duongphuhiep/datastar-helloworld/middlewares"
	datastar "github.com/starfederation/datastar/sdk/go"
)

//go:embed hello-world.html
var helloWorldHTML []byte

// initTracer initializes OpenTelemetry tracing
func initTracer() func() {
	// exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	// if err != nil {
	// 	panic(err)
	// }
	exporter := tracetest.NewNoopExporter()
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("http-api"),
			semconv.ServiceVersionKey.String("1.0.0"),
		)),
	)

	//tp := oteltracenoop.NewTracerProvider()
	otel.SetTracerProvider(tp)

	return func() {
		tp.Shutdown(context.Background())
	}
}

func initLogger(ctx context.Context) func() {
	//Set up OTLP exporter
	exporter, err := otlploggrpc.New(ctx,
		otlploggrpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}

	processor := otlogsdk.NewBatchProcessor(exporter)
	provider := otlogsdk.NewLoggerProvider(otlogsdk.WithProcessor(processor))

	// 3. Bridge OpenTelemetry with slog
	otelLogger := otelslog.NewLogger("otel-logger", otelslog.WithLoggerProvider(provider))

	slog.SetDefault(otelLogger)

	return func() {
		if err := provider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}
}

var globalReqId = 0

func main() {
	ctx := context.Background()

	tracerDispose := initTracer()
	defer tracerDispose()

	loggerDispose := initLogger(ctx)
	defer loggerDispose()

	mux := http.NewServeMux()

	type Store struct {
		Myinput string `json:"myinput"` // delay in milliseconds between each character of the message.
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(helloWorldHTML)
	})

	mux.HandleFunc("/actions/setinput", func(w http.ResponseWriter, r *http.Request) {
		store := &Store{}
		if err := datastar.ReadSignals(r, store); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		globalReqId++
		reqId := strconv.Itoa(globalReqId)
		slog.InfoContext(r.Context(), "Received input from client "+reqId, "myinput", store.Myinput)
		sse := datastar.NewSSE(w, r)
		for range 5 {
			slog.InfoContext(r.Context(), "Send response for "+reqId)
			greetMsg := fmt.Sprintf("%s - time on server is %s", reqId, time.Now().Format(time.RFC3339))
			err := sse.MergeSignals([]byte(`{myinput: '` + greetMsg + `'}`))
			if err != nil {
				panic(err)
			}
			time.Sleep(2 * time.Second)
		}
	})

	// Wrap the mux with the logging middleware
	loggedMux := m.LoggingMiddleware(mux)

	slog.Info("Starting server on :8080")
	http.ListenAndServe(":8080", loggedMux)
}
