package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode) // untuk menentukan log mode (default: Debug)
	os.Exit(m.Run())
}
