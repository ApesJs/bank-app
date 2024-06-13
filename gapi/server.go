package gapi

import (
	db "github.com/ApesJs/bank-app/db/sqlc"
	"github.com/ApesJs/bank-app/pb"
	"github.com/ApesJs/bank-app/util"
)

type Server struct {
	pb.UnimplementedBankAppServer
	store  *db.Store
	config util.Config
}

func NewServer(config util.Config, store *db.Store) *Server {
	server := &Server{
		store:  store,
		config: config,
	}

	return server
}
