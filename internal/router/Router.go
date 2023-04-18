package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	return router
}

func RunApp(port string, router *gin.Engine) error {
	err := router.Run(":" + port)
	if err != nil {
		return fmt.Errorf("error when try to run server")
	}
	return nil
}
