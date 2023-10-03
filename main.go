package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
  c.String(http.StatusOK, "hello\n")
}

func main() {
  router := gin.Default()
  router.GET("/hello", hello)
  
  router.Run("localhost:3000")
}
