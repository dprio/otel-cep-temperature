package opentelemetry

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	once         sync.Once
	setupErr     error
	shutdownFunc func(context.Context) error
	globalTracer trace.Tracer
)

// Init initializes the OpenTelemetry SDK only once.
// It must be called before using StartSpan.
func Init(ctx context.Context, serviceName string) error {
	once.Do(func() {
		var shutdown func(context.Context) error
		shutdown, setupErr = setupOTelSDK(ctx)
		if setupErr != nil {
			return
		}

		shutdownFunc = shutdown
		globalTracer = otel.Tracer(serviceName)
	})

	return setupErr
}

// Shutdown gracefully flushes providers and shuts down OpenTelemetry.
func Shutdown(ctx context.Context) error {
	if shutdownFunc == nil {
		return nil
	}
	return shutdownFunc(ctx)
}

// StartSpan starts a span using the global tracer initialized in Init().
func StartSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if globalTracer == nil {
		panic("opentelemetry.Init must be called before StartSpan")
	}

	return globalTracer.Start(ctx, name, opts...)
}
