package create_transaction

import (
	"context"
	"testing"

	"github.com.br/andersonhqds/fc-ms-wallet/internal/entity"
	"github.com.br/andersonhqds/fc-ms-wallet/internal/event"
	"github.com.br/andersonhqds/fc-ms-wallet/internal/usecase/mocks"
	"github.com.br/andersonhqds/fc-ms-wallet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, err := entity.NewClient("John Doe", "j@j")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)
	assert.Nil(t, err)
	assert.NotNil(t, client1)
	assert.NotNil(t, account1)

	client2, err := entity.NewClient("Jane Doe", "jane@jane")
	account2 := entity.NewAccount(client2)
	assert.Nil(t, err)
	assert.NotNil(t, client2)
	assert.NotNil(t, account2)

	mockUow := &mocks.UowMock{}

	mockUow.On("Do", mock.Anything).Return(nil)

	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()
	eventBalance := event.NewBalanceUpdated()

	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, eventTransaction, eventBalance)
	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        500,
	}

	output, err := uc.Execute(ctx, inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
