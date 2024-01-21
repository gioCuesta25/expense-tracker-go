package api

// type createExpenseReq struct {
// 	Description string `json:"description" binding:"required"`
// 	Amount      int64  `json:"amount" binding:"required,min=1"`
// 	CategoryID  int64  `json:"category_id" binding:"required,min=1"`
// }

// func (server *Server) handleCreateEntry(ctx *gin.Context) {
// 	var req createExpenseReq
// 	/*
// 		ShouldBindJSON se usa para decodificar datos JSON y asignarlos
// 		a una estructura de datos go
// 	*/
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.CreateExpenseParams{
// 		Description: req.Description,
// 		Amount:      req.Amount,
// 		CategoryID:  req.CategoryID,
// 	}

// 	newExpense, err := server.database.CreateExpense(ctx, arg)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, newExpense)
// }

// type getExpenseReq struct {
// 	ID int64 `uri:"id" binding:"required,min=1"`
// }

// func (server *Server) handleGetEntry(ctx *gin.Context) {
// 	var req getExpenseReq

// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	expense, err := server.database.GetExpense(ctx, req.ID)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, expense)
// }

// type deleteExpenseReq struct {
// 	ID int64 `uri:"id" binding:"required,min=1"`
// }

// func (server *Server) handleDeleteEntry(ctx *gin.Context) {
// 	var req getExpenseReq

// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	err := server.database.DeleteExpense(ctx, req.ID)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, nil)
// }
