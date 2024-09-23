package create_transaction

import (
	"testing"

	"github.com.br/andersonhqds/fc-ms-wallet/internal/entity"
	"github.com.br/andersonhqds/fc-ms-wallet/internal/event"
	"github.com.br/andersonhqds/fc-ms-wallet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

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

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("FindByID", account1.ID).Return(account1, nil)
	accountGatewayMock.On("FindByID", account2.ID).Return(account2, nil)

	transactionGatewayMock := &TransactionGatewayMock{}
	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()

	uc := NewCreateTransactionUseCase(transactionGatewayMock, accountGatewayMock, dispatcher, event)
	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        500,
	}

	output, err := uc.Execute(inputDto)

	assert.NotEmpty(t, output.ID)
	assert.Nil(t, err)
	accountGatewayMock.AssertExpectations(t)
	transactionGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGatewayMock.AssertNumberOfCalls(t, "Create", 1)
}
