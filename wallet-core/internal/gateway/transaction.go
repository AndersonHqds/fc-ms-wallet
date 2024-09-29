package gateway

import "github.com.br/andersonhqds/fc-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
