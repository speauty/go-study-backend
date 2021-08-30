package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5433/backend?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB,err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("连接数据库失败：", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
