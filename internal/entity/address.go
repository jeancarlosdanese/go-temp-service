package entity

// Address - struct para armazenar os dados de endereço padronizados
type Address struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	Uf         string `json:"uf"`
}
