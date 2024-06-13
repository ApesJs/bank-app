package gapi

import (
	"context"
	db "github.com/ApesJs/bank-app/db/sqlc"
	"github.com/ApesJs/bank-app/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateTransfer(ctx context.Context, req *pb.CreateTransferRequest) (*pb.CreateTransferResult, error) {
	validAccount := server.validAccount(ctx, req.GetFromAccountId(), req.GetToAccountId())
	isBalanceEnough := server.isUserBalanceEnough(ctx, req.GetFromAccountId(), req.GetAmount())

	if validAccount != "" || !isBalanceEnough {
		message := "your balance is not enough"
		if validAccount != "" {
			message = validAccount
		}

		return &pb.CreateTransferResult{
			Result: &pb.CreateTransferResult_ErrorResponse{
				ErrorResponse: &pb.CreateTransferErrorResponse{
					Status:  "failed",
					Message: message,
				},
			},
		}, nil
	}

	arg := db.TransferTxParams{
		FromAccountID: req.GetFromAccountId(),
		ToAccountID:   req.GetToAccountId(),
		Amount:        req.GetAmount(),
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "transfer failed")
	}

	return &pb.CreateTransferResult{
		Result: &pb.CreateTransferResult_Response{
			Response: &pb.CreateTransferResponse{
				Status:        "success",
				TransactionId: result.Transfer.ID,
				Message:       "Transfer successful",
			},
		},
	}, nil
}

func (server *Server) validAccount(ctx context.Context, accountID, toAccountID int64) string {
	if !server.isAccountValid(ctx, accountID) {
		return "account does not exist"
	}

	if !server.isAccountValid(ctx, toAccountID) {
		return "recipient account does not exist"
	}

	return ""
}

func (server *Server) isAccountValid(ctx context.Context, accountID int64) bool {
	_, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		return false
	}
	return true
}

func (server *Server) isUserBalanceEnough(ctx context.Context, accountID, amount int64) bool {
	account, err := server.store.GetAccount(ctx, accountID)

	if err != nil {
		return false
	}

	if account.Balance < amount {
		return false
	}

	return true
}
