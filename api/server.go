package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/gioSmith25/expense-tracker/db/sqlc"
)

/*
Estructura server para interactuar con la base de datos y
crear el router de la api
*/
type Server struct {
	database *db.Queries
	router   *gin.Engine
}

func NewServer(q *db.Queries) *Server {
	server := Server{database: q}
	router := gin.Default()

	// Aquí se añaden las rutas de la api
	router.POST("/expenses", server.handleCreateExpense)
	router.GET("/expenses/:id", server.handleGetExpense)
	router.DELETE("/expenses/:id", server.handleDeleteExpense)

	server.router = router
	return &server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
