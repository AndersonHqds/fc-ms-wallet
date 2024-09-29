package database

import (
	"database/sql"
	"testing"

	"github.com.br/andersonhqds/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDb *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance decimal, created_at date, FOREIGN KEY(client_id) REFERENCES clients(id))")
	s.accountDb = NewAccountDB(db)
	s.client, _ = entity.NewClient("john", "j@j")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
	defer s.db.Close()
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestAccountDB_Save() {
	account := entity.NewAccount(s.client)
	err := s.accountDb.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestAccountDB_FindByID() {
	s.db.Exec("Insert into clients (id, name, email, created_at) values (?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)
	account := entity.NewAccount(s.client)
	err := s.accountDb.Save(account)
	s.Nil(err)
	accountFromDB, err := s.accountDb.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountFromDB.ID)
	s.Equal(account.Client.ID, accountFromDB.Client.ID)
	s.Equal(account.Balance, accountFromDB.Balance)
	s.Equal(account.Client.Name, accountFromDB.Client.Name)
	s.Equal(account.Client.Email, accountFromDB.Client.Email)
}
