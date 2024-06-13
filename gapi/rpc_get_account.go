package gapi

import (
	"context"
	"errors"
	"github.com/ApesJs/bank-app/pb"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResult, error) {
	account, err := server.store.GetAccount(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &pb.GetAccountResult{
				Result: &pb.GetAccountResult_ErrorResponse{
					ErrorResponse: &pb.GetAccountErrorResponse{
						Error: "User not found",
					},
				},
			}, nil
		}
		return nil, status.Errorf(codes.Internal, "get account failed")
	}

	return &pb.GetAccountResult{
		Result: &pb.GetAccountResult_Response{
			Response: &pb.GetAccountResponse{
				Id:      account.ID,
				Balance: account.Balance,
			},
		},
	}, nil
}
