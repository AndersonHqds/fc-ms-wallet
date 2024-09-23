package database

import (
	"database/sql"
	"testing"

	"github.com.br/andersonhqds/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance int, created_at date, FOREIGN KEY(client_id) REFERENCES clients(id))")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date, FOREIGN KEY(account_id_from) REFERENCES accounts(id), FOREIGN KEY(account_id_to) REFERENCES accounts(id))")
	client, err := entity.NewClient("john", "j@j")
	s.Require().NoError(err)
	s.client = client
	client2, err := entity.NewClient("jane", "jj@j")
	s.Require().NoError(err)
	s.client2 = client2

	accountFrom := entity.NewAccount(client)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(client2)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	s.db.Exec("DROP TABLE transactions")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
	defer s.db.Close()
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestTransactionDB_Create() {
	err := s.transactionDB.Create(&entity.Transaction{
		AccountFrom: s.accountFrom,
		AccountTo:   s.accountTo,
		Amount:      100,
	})
	s.Nil(err)
}
