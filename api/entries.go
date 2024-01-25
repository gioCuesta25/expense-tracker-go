package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/gioSmith25/expense-tracker/db/sqlc"
	"github.com/google/uuid"
)

type createEntryReq struct {
	Description string    `json:"description" binding:"required"`
	Amount      int64     `json:"amount" binding:"required,min=1"`
	CategoryID  uuid.UUID `json:"category_id" binding:"required,min=1"`
	TypeID      uuid.UUID `json:"type_id" binding:"required,min=1"`
}

func (server *Server) handleCreateEntry(ctx *gin.Context) {
	var req createEntryReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEntryParams{
		Description: req.Description,
		Amount:      req.Amount,
		CategoryID:  req.CategoryID,
		TypeID:      req.TypeID,
	}

	newEntry, err := server.database.CreateEntry(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newEntry)
}

type getEntryReq struct {
	ID uuid.UUID `uri:"id" binding:"required,min=1"`
}

func (server *Server) handleGetEntry(ctx *gin.Context) {
	var req getEntryReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	entry, err := server.database.GetEntry(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entry)

}

type deleteEntryReq struct {
	ID uuid.UUID `uri:"id" binding:"required,min=1"`
}

func (server *Server) handleDeleteEntry(ctx *gin.Context) {
	var req deleteEntryReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err := server.database.DeleteEntry(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)

}

type listEntriesReq struct {
	Page int32 `form:"page"`
	Size int32 `form:"size"`
}

type listEntriesResponse struct {
	Items []db.Entry `json:"items"`
	Total int        `json:"total"`
	Page  int        `json:"page"`
	Size  int        `json:"size"`
}

func (server *Server) handleListEntries(ctx *gin.Context) {
	var params listEntriesReq

	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListEntriesParams{
		Limit:  params.Size,
		Offset: (params.Page - 1) * params.Size,
	}

	entries, err := server.database.ListEntries(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := listEntriesResponse{
		Items: entries,
		Page:  int(params.Page),
		Size:  int(params.Size),
		Total: len(entries),
	}

	ctx.JSON(http.StatusOK, response)

}
