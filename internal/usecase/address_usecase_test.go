package usecase_test

import (
	"context"
	"testing"

	"github.com/jeancarlosdanese/go-temp-service/internal/entity"
	"github.com/jeancarlosdanese/go-temp-service/internal/usecase"
	"github.com/stretchr/testify/mock"
)

type MockAddressClient struct {
	mock.Mock
}

func (m *MockAddressClient) FetchAddress(ctx context.Context, cep string) (*entity.Address, error) {
	args := m.Called(ctx, cep)
	return args.Get(0).(*entity.Address), args.Error(1)
}

func TestGetAddress(t *testing.T) {
	mockClient := new(MockAddressClient)
	mockClient.On("FetchAddress", mock.Anything, "12345678").Return(&entity.Address{
		Cep:    "12345678",
		Cidade: "São Paulo",
		Uf:     "SP",
	}, nil)

	usecase := usecase.NewAddressUsecase(mockClient)
	address, err := usecase.GetAddress(context.Background(), "12345678")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if address.Cidade != "São Paulo" {
		t.Errorf("Expected city 'São Paulo', got '%s'", address.Cidade)
	}
}
