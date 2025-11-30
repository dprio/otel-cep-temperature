package weatherhandler

import "github.com/dprio/cep-temperature/internal/domain/weather"

type Response struct {
	CelsiusTemp       float64 `json:"temp_c"`
	FahrenheitTemp    float64 `json:"temp_f"`
	KelvinTemperature float64 `json:"temp_k"`
}

func NewResponse(weather *weather.Weather) Response {
	if weather == nil {
		return Response{}
	}

	return Response{
		CelsiusTemp:       weather.Temperature.C,
		FahrenheitTemp:    weather.Temperature.F,
		KelvinTemperature: weather.Temperature.K,
	}
}
