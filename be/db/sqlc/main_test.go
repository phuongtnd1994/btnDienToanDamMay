package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/phuongtnd/test01/util"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB
var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	testDb, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDb)
	testStore = NewStore(testDb)

	os.Exit(m.Run())
}
