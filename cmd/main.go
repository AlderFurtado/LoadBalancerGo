package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	apiName := os.Getenv("API_NAME")
	if apiName == "" {
		apiName = "API_UNKNOW"
	}

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"result": apiName,
		})
	})

	router.GET("/health", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"result": "ok",
		})
	})

	router.Run(":3000")
}
