package db

import (
	"log"
	"os"
	"testing"

	"github.com/blog/config"
	"github.com/blog/internal/storage/repo"
	"github.com/blog/pkg/db"
)

var testMain repo.BlogService

func TestMain(t *testing.M) {
	cfg := config.Load("../../../")

	conn, err := db.ConnectToDb(cfg)

	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}
	testMain = NewStorage(conn)
	os.Exit(t.Run())
}
