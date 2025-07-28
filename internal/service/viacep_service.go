package service

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Erro       string `json:"erro,omitempty"`
}

type ViaCEPServiceInterface interface {
	GetViaCEPInfo(cep string) (*ViaCEPResponse, error)
}
