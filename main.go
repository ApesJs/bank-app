package main

import (
	"context"
	db "github.com/ApesJs/bank-app/db/sqlc"
	"github.com/ApesJs/bank-app/gapi"
	"github.com/ApesJs/bank-app/pb"
	"github.com/ApesJs/bank-app/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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
	defer conn.Close()

	store := db.NewStore(conn)
	server := gapi.NewServer(config, store)

	grpcServer := grpc.NewServer()
	pb.RegisterBankAppServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}

}
