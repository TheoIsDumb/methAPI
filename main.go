package main

import (
	"net/http"
	"time"

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
  name := c.Param("name")
  c.String(http.StatusOK, "goodbye %s!\n", name)
}

func datecalc(c *gin.Context) {
  dateLayout := "2006-01-02"
  first := c.Query("f")
  last := c.DefaultQuery("l", time.Now().Format("2006-01-02"))
  firstDate, _ := time.Parse(dateLayout, first)
  secondDate, _ := time.Parse(dateLayout, last)
  difference := firstDate.Sub(secondDate)

  c.String(http.StatusOK, "%v\n", difference.Abs().Hours()/24)
}

func main() {
  router := gin.Default()
  router.GET("/", root)
  router.GET("/hello", hello)
  router.GET("/goodbye/:name", goodbye)
  router.GET("/datecalc", datecalc)
  
  router.Run("localhost:3000")
}