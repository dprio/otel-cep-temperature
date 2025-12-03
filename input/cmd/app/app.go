package app

import (
	"context"
	"errors"
	"os"
	"os/signal"

	"github.com/dprio/otel-cep-temperature/input/internal/gateway"
	"github.com/dprio/otel-cep-temperature/input/internal/infrastructure/config"
	"github.com/dprio/otel-cep-temperature/input/internal/infrastructure/httpclient"
	"github.com/dprio/otel-cep-temperature/input/internal/infrastructure/web/handlers"
	"github.com/dprio/otel-cep-temperature/input/internal/infrastructure/web/webserver"
	"github.com/dprio/otel-cep-temperature/input/internal/usecases"
	"github.com/dprio/otel-cep-temperature/input/pkg/opentelemetry"
)

type App struct {
	webServer *webserver.WebServer
}

func New() *App {
	cfg := config.New()

	httpClients := httpclient.New()

	gateways := gateway.New(*httpClients)

	useCases := usecases.New(gateways)

	hdls := handlers.New(useCases)

	webServer := createWebServer(cfg, hdls)

	return &App{
		webServer: webServer,
	}
}

func createWebServer(cfg *config.Config, hdls *handlers.Handlers) *webserver.WebServer {
	webServer := webserver.New(cfg.Web)

	webServer.AddHandler("POST", "/addresses/weather", hdls.WeatherHandler.GetLocationTemperature)

	return webServer
}

func (a *App) Start() error {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	otelShutdown, err := opentelemetry.SetupOTelSDK(ctx)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	return a.webServer.Start()
}
