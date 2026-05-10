package viacephttpclient

type Response struct {
	ZipCode string `json:"cep"`
	City    string `json:"localidade"`
	Erro    string `json:"erro"`
}

func (r *Response) isNotFound() bool {
	return r.Erro == "true"
}
