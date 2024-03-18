package main

import (
	"context"
	"github.com/ApesJs/bank-app/api"
	db "github.com/ApesJs/bank-app/db/sqlc"
	"github.com/ApesJs/bank-app/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}
