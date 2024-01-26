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

	router.POST("/entry", server.handleCreateEntry)
	router.GET("/entry/:id", server.handleGetEntry)
	router.GET("/entry", server.handleListEntries)
	router.DELETE("/entry/:id", server.handleDeleteEntry)

	router.GET("/category", server.handleListCategories)

	server.router = router
	return &server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
