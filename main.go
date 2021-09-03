package main

import (
	"database/sql"
	api "github.com/speauty/backend/src/api"
	db "github.com/speauty/backend/src/db/sqlc"
	util "github.com/speauty/backend/src/util"
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
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("创建服务失败:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("启动服务失败:", err)
	}
}
