package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    

    

    
}

func SetupRouter() *gin.Engine {
  app := gin.Default();

  app.Use(cors.New(cors.Config{
    AllowOrigins: []string{"*"},
  }))

}