package weatherapihttpclient

type (
	Response struct {
		Current Current `json:"current"`
	}

	Current struct {
		TempC float64 `json:"temp_c"`
	}
)
