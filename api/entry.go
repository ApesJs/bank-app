package api

import (
	"errors"
	db "github.com/ApesJs/bank-app/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
)

type getEntriesRequest struct {
	ID       int64 `form:"id" binding:"required,min=1"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type entryResponse struct {
	TransferID int64  `json:"transaction_id"`
	CreatedAt  string `json:"date"`
	Amount     int64  `json:"amount"`
	Type       string `json:"type"`
}

func (server *Server) getEntries(ctx *gin.Context) {
	var req getEntriesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListEntriesParams{
		AccountID: req.ID,
		Limit:     req.PageSize,
		Offset:    (req.PageID - 1) * req.PageSize,
	}

	entries, err := server.store.ListEntries(ctx, arg)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var entryResponses []entryResponse
	for _, entry := range entries {
		entryResponses = append(entryResponses, entryResponse{
			TransferID: entry.TransferID,
			CreatedAt:  entry.CreatedAt.Time.Format("2006-01-02"),
			Amount:     entry.Amount,
			Type:       entry.Type,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"account_id":   req.ID,
		"transactions": entryResponses,
	})
}
