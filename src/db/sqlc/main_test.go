package db

import (
	"database/sql"
	util2 "github.com/speauty/backend/src/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util2.LoadConfig("../../..")
	if err != nil {
		log.Fatal("载入配置失败: ", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("连接数据库失败：", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
