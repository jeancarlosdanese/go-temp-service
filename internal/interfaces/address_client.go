// internal/interfaces/address_client.go

package interfaces

import (
	"context"

	"github.com/jeancarlosdanese/go-temp-service/internal/entity"
)

type AddressClientInterface interface {
	FetchAddress(ctx context.Context, cep string) (*entity.Address, error)
}
