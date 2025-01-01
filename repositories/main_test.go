package repositories

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/toko_buku_online_nextjs?sslmode=disable"
)

var UsersRepoTest *UsersRepo
var BooksRepoTest *BooksRepo

func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatalln(err.Error())
	}

	BooksRepoTest = NewBooksRepo(conn)
	UsersRepoTest = NewUsersRepo(conn)

	os.Exit(m.Run())

}
