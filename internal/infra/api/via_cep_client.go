// internal/infra/api/via_cep_client.go

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jeancarlosdanese/go-temp-service/internal/entity"
	"github.com/jeancarlosdanese/go-temp-service/internal/interfaces"
)

type ViaCepClient struct {
	BaseURL string
}

type ViaCEP struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func (v *ViaCepClient) FetchAddress(ctx context.Context, cep string) (*entity.Address, error) {
	// Use a BaseURL se estiver definida, senão use a URL padrão
	url := fmt.Sprintf("%s/ws/%s/json/", v.getBaseURL(), cep)
	log.Printf("Buscando endereço de %s: %s\n", entity.ViaCepName, url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching address: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error from %s: CEP not found or other error", entity.ViaCepName)
	}

	var viaCep ViaCEP
	if err := json.NewDecoder(resp.Body).Decode(&viaCep); err != nil {
		return nil, fmt.Errorf("error decoding ViaCEP response: %w", err)
	}

	if viaCep.Cep == "" {
		return nil, fmt.Errorf("error from %s: CEP not found", entity.ViaCepName)
	}

	addr := &entity.Address{
		Cep:        viaCep.Cep,
		Logradouro: viaCep.Logradouro,
		Bairro:     viaCep.Bairro,
		Cidade:     viaCep.Localidade,
		Uf:         viaCep.Uf,
	}
	return addr, nil
}

func (v *ViaCepClient) getBaseURL() string {
	if v.BaseURL != "" {
		return v.BaseURL
	}
	return "https://viacep.com.br"
}

// Certificando que ViaCepClient implementa AddressClientInterface
var _ interfaces.AddressClientInterface = (*ViaCepClient)(nil)
