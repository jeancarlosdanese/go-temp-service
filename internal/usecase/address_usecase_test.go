package usecase

import (
	"context"
	"testing"

	"github.com/jeancarlosdanese/go-temp-service/internal/entity"
	"github.com/stretchr/testify/assert"
)

type mockAddressClient struct{}

func (m *mockAddressClient) FetchAddress(ctx context.Context, cep string) (*entity.Address, error) {
	return &entity.Address{
		Cep:        "01001-000",
		Logradouro: "Praça da Sé",
		Bairro:     "Sé",
		Cidade:     "São Paulo",
		Uf:         "SP",
	}, nil
}

func TestAddressUsecase_GetAddress(t *testing.T) {
	ctx := context.Background()
	mockClient := &mockAddressClient{}
	usecase := NewAddressUsecase(mockClient)

	address, err := usecase.GetAddress(ctx, "01001000")

	assert.NoError(t, err)
	assert.Equal(t, &entity.Address{
		Cep:        "01001-000",
		Logradouro: "Praça da Sé",
		Bairro:     "Sé",
		Cidade:     "São Paulo",
		Uf:         "SP",
	}, address)
}
