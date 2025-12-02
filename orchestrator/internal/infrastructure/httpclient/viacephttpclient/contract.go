package viacephttpclient

type Response struct {
	ZipCode string `json:"cep"`
	City    string `json:"localidade"`
}
