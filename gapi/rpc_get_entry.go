package gapi

import (
	"context"
	"errors"
	db "github.com/ApesJs/bank-app/db/sqlc"
	"github.com/ApesJs/bank-app/pb"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetEntry(ctx context.Context, req *pb.GetEntryRequest) (*pb.GetEntryResult, error) {
	arg := db.ListEntriesParams{
		AccountID: req.GetAccountId(),
		Limit:     req.GetPageSize(),
		Offset:    (req.GetPageId() - 1) * req.GetPageSize(),
	}

	entries, err := server.store.ListEntries(ctx, arg)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &pb.GetEntryResult{
				Result: &pb.GetEntryResult_ErrorResponse{
					ErrorResponse: &pb.GetEntryErrorResponse{
						Error: "Data entry not found",
					},
				},
			}, nil
		}

		return nil, status.Errorf(codes.Internal, "get entry failed")
	}

	var transactionResponses []*pb.TransactionResponse
	for _, entry := range entries {
		transactionResponses = append(transactionResponses, &pb.TransactionResponse{
			TransactionId: entry.TransferID,
			Date:          entry.CreatedAt.Time.Format("2006-01-02"),
			Amount:        entry.Amount,
			Type:          entry.Type,
		})
	}

	return &pb.GetEntryResult{
		Result: &pb.GetEntryResult_Response{
			Response: &pb.GetEntryResponse{
				AccountId:    req.GetAccountId(),
				Transactions: transactionResponses,
			},
		},
	}, nil
}
