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

type BrasilApiClient struct {
	BaseURL string
}

type BrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func (v *BrasilApiClient) FetchAddress(ctx context.Context, cep string) (*entity.Address, error) {
	url := fmt.Sprintf("%s/api/cep/v1/%s", v.getBaseURL(), cep)
	log.Printf("Buscando endereço de %s: %s\n", entity.BrasilApiName, url)

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
		return nil, fmt.Errorf("error from %s: CEP not found or other error", entity.BrasilApiName)
	}

	var brasilApi BrasilAPI
	if err := json.NewDecoder(resp.Body).Decode(&brasilApi); err != nil {
		return nil, fmt.Errorf("error decoding BrasilAPI response: %w", err)
	}

	addr := &entity.Address{
		Cep:        brasilApi.Cep,
		Logradouro: brasilApi.Street,
		Bairro:     brasilApi.Neighborhood,
		Cidade:     brasilApi.City,
		Uf:         brasilApi.State,
	}
	return addr, nil
}

func (v *BrasilApiClient) getBaseURL() string {
	if v.BaseURL != "" {
		return v.BaseURL
	}
	return "https://brasilapi.com.br"
}

// Certificando que BrasilApiClient implementa AddressClientInterface
var _ interfaces.AddressClientInterface = (*BrasilApiClient)(nil)
