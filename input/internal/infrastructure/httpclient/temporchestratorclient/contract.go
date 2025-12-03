package temporchestratorclient

type Response struct {
	City              string  `json:"city"`
	CelsiusTemp       float64 `json:"temp_c"`
	FahrenheitTemp    float64 `json:"temp_f"`
	KelvinTemperature float64 `json:"temp_k"`
}
