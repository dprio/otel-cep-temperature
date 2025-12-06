package weatherhandler

import (
	"encoding/json"
	"net/http"

	"github.com/dprio/otel-cep-temperature/orchestrator/internal/usecases/gettemperaturebyzipcode"
	"github.com/dprio/otel-cep-temperature/orchestrator/pkg/opentelemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
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
	ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))

	ctx, span := opentelemetry.StartSpan(ctx, "WeatherHandler.GetLocationTemperature")
	defer span.End()

	zipCode := r.PathValue("ZIP_CODE")

	span.SetAttributes(attribute.String("cep", zipCode))
	output, err := h.getTemperatureByZipCodeUseCase.Execute(ctx, zipCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
