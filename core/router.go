package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	return router
}
