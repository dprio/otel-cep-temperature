package app

import (
	"github.com/dprio/otel-cep-temperature/orchestrator/internal/gateway"
	"github.com/dprio/otel-cep-temperature/orchestrator/internal/infrastructure/config"
	"github.com/dprio/otel-cep-temperature/orchestrator/internal/infrastructure/httpclient"
	"github.com/dprio/otel-cep-temperature/orchestrator/internal/infrastructure/web/handlers"
	"github.com/dprio/otel-cep-temperature/orchestrator/internal/infrastructure/web/webserver"
	"github.com/dprio/otel-cep-temperature/orchestrator/internal/usecases"
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

	webServer.AddHandler("GET", "/addresses/{ZIP_CODE}/weather/temperature", hdls.WeatherHandler.GetLocationTemperature)

	return webServer
}

func (a *App) Start() error {
	return a.webServer.Start()
}
