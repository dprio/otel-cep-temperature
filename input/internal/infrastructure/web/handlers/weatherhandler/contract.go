package weatherhandler

import (
	"github.com/dprio/otel-cep-temperature/input/internal/usecases/gettemperaturebyzipcode"
)

type (
	Request struct {
		ZipCode string `json:"cep"`
	}

	Response struct {
		City              string  `json:"city"`
		CelsiusTemp       float64 `json:"temp_c"`
		FahrenheitTemp    float64 `json:"temp_f"`
		KelvinTemperature float64 `json:"temp_k"`
	}
)

func NewResponse(out *gettemperaturebyzipcode.Output) Response {
	if out == nil {
		return Response{}
	}

	return Response{
		City:              out.Weather.City,
		CelsiusTemp:       out.Weather.C,
		FahrenheitTemp:    out.Weather.F,
		KelvinTemperature: out.Weather.K,
	}
}
