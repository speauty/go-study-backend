package main

import (
	"database/sql"
	"github.com/backend/api"
	db "github.com/backend/db/sqlc"
	"github.com/backend/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("载入配置失败: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("连接数据库失败: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("启动服务失败:", err)
	}
}
