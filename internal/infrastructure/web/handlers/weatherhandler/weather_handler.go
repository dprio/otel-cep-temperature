package weatherhandler

import (
	"encoding/json"
	"net/http"

	"github.com/dprio/otel-cep-temperature/internal/usecases/gettemperaturebyzipcode"
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

	zipCode := r.PathValue("ZIP_CODE")

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
