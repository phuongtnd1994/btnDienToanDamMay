package main

import (
	"database/sql"
	"log"

	db "github.com/phuongtnd/test01/db/sqlc"

	"github.com/phuongtnd/test01/util"

	"github.com/phuongtnd/test01/api"

	_ "github.com/jackc/pgx/v5"
)

func main() {
	log.Printf("Finish 11")
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	// redisOpt := asynq.RedisClientOpt{
	// 	Addr: config.RedisAddress,
	// }

	// taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}

// func runTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) {
// 	taskProcessor := worker.NewRedisTaskDistributor(redisOpt, store)
// }
