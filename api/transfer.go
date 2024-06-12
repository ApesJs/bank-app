package api

import (
	db "github.com/ApesJs/bank-app/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type transferRequest struct {
	FromAccountID int64 `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64 `json:"to_account_id" binding:"required,min=1"`
	Amount        int64 `json:"amount" binding:"required,gt=0"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	validAccount := server.validAccount(ctx, req.FromAccountID, req.ToAccountID)
	isBalanceEnough := server.isUserBalanceEnough(ctx, req.FromAccountID, req.Amount)

	if validAccount != "" || !isBalanceEnough {
		message := "your balance is not enough"
		if validAccount != "" {
			message = validAccount
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": message,
		})
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"transfer_id": result.Transfer.ID,
		"message":     "Transfer Successful",
	})
}

func (server *Server) validAccount(ctx *gin.Context, accountID, toAccountID int64) string {
	if !server.isAccountValid(ctx, accountID) {
		return "account does not exist"
	}

	if !server.isAccountValid(ctx, toAccountID) {
		return "recipient account does not exist"
	}

	return ""
}

func (server *Server) isAccountValid(ctx *gin.Context, accountID int64) bool {
	_, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		return false
	}
	return true
}

func (server *Server) isUserBalanceEnough(ctx *gin.Context, accountID, amount int64) bool {
	account, err := server.store.GetAccount(ctx, accountID)

	if err != nil {
		return false
	}

	if account.Balance < amount {
		return false
	}

	return true
}
