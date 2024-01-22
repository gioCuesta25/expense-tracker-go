package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/gioSmith25/expense-tracker/db/sqlc"
)

type listCategoriesResponse struct {
	Items []db.Category `json:"items"`
	Total int           `json:"total"`
}

func (server *Server) handleListCategories(ctx *gin.Context) {
	categories, err := server.database.ListCategories(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	response := listCategoriesResponse{
		Items: categories,
		Total: len(categories),
	}

	ctx.JSON(http.StatusOK, response)
}
