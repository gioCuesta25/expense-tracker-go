package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/gioSmith25/expense-tracker/db/sqlc"
	"github.com/gioSmith25/expense-tracker/token"
	"github.com/gioSmith25/expense-tracker/utils"
)

/*
Estructura server para interactuar con la base de datos y
crear el router de la api
*/
type Server struct {
	config     utils.Config
	database   *db.Queries
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, q *db.Queries) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.JWTSecretKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := Server{
		database:   q,
		config:     config,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return &server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/entry", server.handleCreateEntry)
	router.GET("/entry/:id", server.handleGetEntry)
	router.GET("/entry", server.handleListEntries)
	router.DELETE("/entry/:id", server.handleDeleteEntry)

	router.GET("/category", server.handleListCategories)

	router.POST("/users", server.handleCreateUser)
	router.GET("/users/:id", server.handleGetUser)

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
