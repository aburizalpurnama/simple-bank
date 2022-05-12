package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.getListAccounts)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

type listAccounstRequest struct {
	Limit  int32 `json:"limit" binding:"required,numeric,min=1"`
	Offset int32 `json:"offset" binding:"required,numeric,min=1"`
}

func (s *Server) getListAccounts(ctx *gin.Context) {
	var req listAccounstRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.ListAccountsParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	listAcount, err := s.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listAcount)
}

func (s *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := s.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
