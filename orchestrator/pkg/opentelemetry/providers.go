package opentelemetry

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTracerProvider(serviceName string) (*sdktrace.TracerProvider, error) {
	ctx := context.Background()

	rsrc, err := resource.New(
		ctx,
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
		),
	)
	if err != nil {
		return nil, err

	}

	conn, err := grpc.NewClient("otel-collector:4317",
		grpc.WithIdleTimeout(time.Second),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)

	return sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithResource(rsrc),
	), nil
}

func newMeterProvider() (*metric.MeterProvider, error) {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	reader := metric.NewPeriodicReader(metricExporter, metric.WithInterval(3*time.Second))
	return metric.NewMeterProvider(metric.WithReader(reader)), nil
}

func newLoggerProvider() (*log.LoggerProvider, error) {
	exporter, err := stdoutlog.New()
	if err != nil {
		return nil, err
	}

	return log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(exporter)),
	), nil
}
