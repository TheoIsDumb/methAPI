package main

import (
	"net/http"
	"strconv"
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

  if c.Query("f") == "" {
    c.String(http.StatusBadRequest, "%s\n", "Date(s) not provided.")
  } else {
    first := c.Query("f")
    last := c.DefaultQuery("l", time.Now().Format(dateLayout))
    firstDate, _ := time.Parse(dateLayout, first)
    secondDate, _ := time.Parse(dateLayout, last)
    difference := firstDate.Sub(secondDate)

    c.String(http.StatusOK, "%v\n", difference.Abs().Hours()/24)
  }
}

func brrrcalc(c *gin.Context) {
  if c.Query("d") == "" || c.Query("s") == "" {
    c.String(http.StatusBadRequest, "%s\n", "Duration/Speed not given.")
  } else {
    duration, _ := strconv.ParseFloat(c.Query("d"), 32)
    speed, _ := strconv.ParseFloat(c.Query("s"), 32)

    c.String(http.StatusOK, "%v\n", duration/speed)
  }
}

func main() {
  router := gin.Default()
  router.GET("/", root)
  router.GET("/hello", hello)
  router.GET("/goodbye/:name", goodbye)
  router.GET("/datecalc", datecalc)
  router.GET("/brrrcalc", brrrcalc)
  
  router.Run("localhost:3000")
}