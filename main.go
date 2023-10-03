package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func root(c *gin.Context) {
  c.String(http.StatusOK, "welcome!\n")
}

func hello(c *gin.Context) {
  name := c.DefaultQuery("name", "friend")
  c.String(http.StatusOK, "hello %s!\n", name)
}

func goodbye(c *gin.Context) {
  name, err := c.Param("name")
  if err != nil {
    name = "friend"
  }
  c.String(http.StatusOK, "goodbye %s!\n", name)
}

func main() {
  router := gin.Default()
  router.GET("/", root)
  router.GET("/hello", hello)
  router.GET("/goodbye/:name", goodbye)
  
  router.Run("localhost:3000")
}
