package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var port string

func main() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := gin.Default()

	router.GET("/health", health)
	router.GET("/argo", argo)

	log.Fatal(router.Run(fmt.Sprintf("0.0.0.0:%s", port)))
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"message": "api is working",
	})
}

func argo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"message": "our CI/CD pipeline is working",
		"timestamp": time.Now(),
	})
}
