package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/token"
	"github.com/techschool/simplebank/util"
)

type Server struct {
	store  db.Store
	router *gin.Engine
	maker  token.Maker
	config util.Config
}

func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{store: store, maker: tokenMaker, config: config}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupServer()

	return server, nil
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

// buatkan package khusus untuk response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupServer() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.getListAccounts)
	router.GET("/accounts/:id", server.getAccount) // tanda : untuk menandakan bahwa itu adalah url parameter

	router.POST("/transfers", server.createTransfer)

	server.router = router
}
