package weatherhandler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dprio/otel-cep-temperature/input/internal/infrastructure/httpclient/temporchestratorclient"
	"github.com/dprio/otel-cep-temperature/input/internal/usecases/gettemperaturebyzipcode"
	"github.com/dprio/otel-cep-temperature/input/pkg/opentelemetry"
	"go.opentelemetry.io/otel/attribute"
)

type WeatherHandler struct {
	getTemperatureByZipCodeUseCase gettemperaturebyzipcode.UseCase
}

func New(getTemperatureByZipCodeUseCase gettemperaturebyzipcode.UseCase) *WeatherHandler {
	return &WeatherHandler{
		getTemperatureByZipCodeUseCase: getTemperatureByZipCodeUseCase,
	}
}

func (h *WeatherHandler) GetLocationTemperature(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx, span := opentelemetry.StartSpan(ctx, "WeatherHandler.GetLocationTemperature")
	defer span.End()

	println("Chamando o trem !")

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	span.SetAttributes(attribute.String("cep", req.ZipCode))

	output, err := h.getTemperatureByZipCodeUseCase.Execute(ctx, req.ZipCode)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, temporchestratorclient.ErrNotFound) {
			status = http.StatusNotFound
		}

		http.Error(w, err.Error(), status)
		return
	}

	response := NewResponse(output)
	w.Header().Add("content-type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
