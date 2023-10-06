package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request) {
  var name string

  if r.URL.Query().Has("name") {
    name = r.URL.Query().Get("name")
  } else {
    name = "friend"
  }

  response := fmt.Sprintf("hello %s!\n", name)
	io.WriteString(w, response)
}

func goodbye(w http.ResponseWriter, r *http.Request) {
  var name string

  if r.URL.Query().Has("name") {
    name = r.URL.Query().Get("name")
  } else {
    name = "friend"
  }

  response := fmt.Sprintf("goodbye %s!\n", name)
	io.WriteString(w, response)
}

func datecalc(w http.ResponseWriter, r *http.Request) {
  dateLayout := "2006-01-02"
  var first string
  var second string

  if r.URL.Query().Has("l") {
    second = r.URL.Query().Get("l")
  } else {
    second = time.Now().Format(dateLayout)
  }

  if r.URL.Query().Has("f") {
    first = r.URL.Query().Get("f")
    firstDate, _ := time.Parse(dateLayout, first)
    secondDate, _ := time.Parse(dateLayout, second)
    difference := firstDate.Sub(secondDate)

    io.WriteString(w, strconv.FormatFloat(difference.Abs().Hours()/24, 'f', 6, 64) + "\n")
  } else {
    io.WriteString(w, "exit")
  }
}

// func brrrcalc(c *gin.Context) {
//   if c.Query("d") == "" || c.Query("s") == "" {
//     c.String(http.StatusBadRequest, "%s\n", "Duration/Speed not given.")
//   } else {
//     duration, _ := strconv.ParseFloat(c.Query("d"), 32)
//     speed, _ := strconv.ParseFloat(c.Query("s"), 32)

//     c.String(http.StatusOK, "%v\n", duration/speed)
//   }
// }

func main() {
  http.HandleFunc("/", root)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye", goodbye)
	http.HandleFunc("/datecalc", datecalc)

  fmt.Println("server started on :3000!")

	err := http.ListenAndServe(":3000", nil)
  if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed!\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}