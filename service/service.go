package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewServer configures and returns a Server.
func NewServer() *gin.Engine {
	// set release mode to turn off debug mode
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	initRoutes(router)

	return router
}

// initRoutes mount the test router
func initRoutes(router *gin.Engine) {
	router.GET("/hello/:id", testHandler)
}

func testHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	// Reply hello message to user
	ctx.String(http.StatusOK, "Hello "+id)
}
