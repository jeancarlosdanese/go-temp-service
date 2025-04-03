// internal/usecase/address_usecase.go

package usecase

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jeancarlosdanese/go-temp-service/internal/entity"
	"github.com/jeancarlosdanese/go-temp-service/internal/interfaces"
)

type AddressUsecase struct {
	clients []interfaces.AddressClientInterface
}

func NewAddressUsecase(clients ...interfaces.AddressClientInterface) *AddressUsecase {
	return &AddressUsecase{clients: clients}
}

func (u *AddressUsecase) GetAddress(ctx context.Context, cep string) (*entity.Address, error) {
	resultCh := make(chan *entity.Address, len(u.clients))
	var wg sync.WaitGroup

	for _, client := range u.clients {
		wg.Add(1)
		go func(client interfaces.AddressClientInterface) {
			defer wg.Done()
			address, err := client.FetchAddress(ctx, cep)
			if err != nil {
				log.Printf("Client failed: %v\n", err)
				return
			}
			select {
			case resultCh <- address:
			case <-ctx.Done():
			}
		}(client)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	select {
	case address := <-resultCh:
		if address != nil {
			return address, nil
		}
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout reached while fetching address for CEP %s", cep)
	}

	return nil, fmt.Errorf("could not find address for CEP %s", cep)
}
