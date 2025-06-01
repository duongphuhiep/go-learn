package toolspack

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	otlogsdk "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// initTracer initializes OpenTelemetry tracing
func InitTracer() func() {
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

func InitOtelLogger(ctx context.Context) func() {
	// Set up OTLP exporter
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
