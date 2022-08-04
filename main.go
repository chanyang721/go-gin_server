package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default();

    r.GET("/ping", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
        "message": "pong",
      })
    })
    r.Run()
  }


// # run example.go and visit 0.0.0.0:8080/ping (for windows "localhost:8080/ping") on browser
// $ go run example.go