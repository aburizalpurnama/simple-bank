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
	router.GET("/accounts/:id", server.getAccount) // tanda : untuk menandakan bahwa itu adalah url parameter

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

// buatkan package khusus untuk request model
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

type listAccounstRequest struct {
	PageSize int32 `json:"page_size" binding:"required,min=5,max=10"`
	PageId   int32 `json:"page_id" binding:"required,min=1"`
}

type getAccountRequst struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequst

	if err := ctx.ShouldBindUri(&req); err != nil { // karena menggunakan url parameter, menggunakan ctx.ShouldBindUri()
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	account, err := s.store.GetAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (s *Server) getListAccounts(ctx *gin.Context) {
	var req listAccounstRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
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

// buatkan package khusus untuk response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
