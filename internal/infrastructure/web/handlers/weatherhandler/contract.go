package weatherhandler

import (
	"github.com/dprio/otel-cep-temperature/internal/usecases/gettemperaturebyzipcode"
)

type Response struct {
	City              string  `json:"city"`
	CelsiusTemp       float64 `json:"temp_c"`
	FahrenheitTemp    float64 `json:"temp_f"`
	KelvinTemperature float64 `json:"temp_k"`
}

func NewResponse(out *gettemperaturebyzipcode.Output) Response {
	if out == nil {
		return Response{}
	}

	return Response{
		City:              out.Address.City,
		CelsiusTemp:       out.Weather.Temperature.C,
		FahrenheitTemp:    out.Weather.Temperature.F,
		KelvinTemperature: out.Weather.Temperature.K,
	}
}
