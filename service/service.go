package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewServer configures and returns a Server.
func NewServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	initRoutes(router)

	return router
}

func initRoutes(router *gin.Engine) {
	router.GET("/hello/:id", testHandler)
}

func testHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "Hello "+id)
}
