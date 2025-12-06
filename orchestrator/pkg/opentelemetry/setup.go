package opentelemetry

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
)

// setupOTelSDK bootstraps the entire OpenTelemetry pipeline.
func setupOTelSDK(ctx context.Context, serviceName string) (func(context.Context) error, error) {
	var shutdownFuncs []func(context.Context) error
	var err error

	// Shutdown executor
	shutdown := func(ctx context.Context) error {
		var e error
		for _, fn := range shutdownFuncs {
			e = errors.Join(e, fn(ctx))
		}
		return e
	}

	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	// Propagators
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// Tracing
	tracerProvider, err := newTracerProvider(serviceName)
	if err != nil {
		handleErr(err)
		return shutdown, err
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	// Metrics
	meterProvider, err := newMeterProvider()
	if err != nil {
		handleErr(err)
		return shutdown, err
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	// Logs
	loggerProvider, err := newLoggerProvider()
	if err != nil {
		handleErr(err)
		return shutdown, err
	}
	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	global.SetLoggerProvider(loggerProvider)

	return shutdown, err
}
