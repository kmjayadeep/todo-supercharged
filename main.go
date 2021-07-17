package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
  app := gin.Default()

  app.GET("/v1", func(c *gin.Context){
    c.JSON(200, gin.H{
      "message": "Welcome to Todo supercharged",
    })
  })

  app.Run()

}
