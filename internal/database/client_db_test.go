package database

import (
	"database/sql"
	"testing"

	"github.com.br/andersonhqds/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Require().NoError(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	s.db.Exec("DROP TABLE clients")
	defer s.db.Close()
}

func TestClientDbTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestClientDB_Save() {
	client, _ := entity.NewClient("john", "j@j")
	err := s.clientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDBTestSuite) TestClientDB_Get() {
	client, _ := entity.NewClient("john", "j@j")
	s.clientDB.Save(client)
	clientFromDB, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientFromDB.ID)
	s.Equal(client.Name, clientFromDB.Name)
	s.Equal(client.Email, clientFromDB.Email)
}
