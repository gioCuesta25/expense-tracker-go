package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/gioSmith25/expense-tracker/db/sqlc"
)

type createExpenseReq struct {
	Description string `json:"description" binding:"required"`
	Amount      int64  `json:"amount" binding:"required,min=1"`
	CategoryID  int64  `json:"category_id" binding:"required,min=1"`
}

func (server *Server) handleCreateExpense(ctx *gin.Context) {
	var req createExpenseReq

	/*
		ShouldBindJSON se usa para decodificar datos JSON y asignarlos
		a una estructura de datos go
	*/
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateExpenseParams{
		Description: req.Description,
		Amount:      req.Amount,
		CategoryID:  req.CategoryID,
	}

	newExpense, err := server.database.CreateExpense(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newExpense)

}
