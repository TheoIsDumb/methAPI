package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request) {
  var name string

  if r.URL.Query().Has("name") {
    name = r.URL.Query().Get("name")
  } else {
    name = "friend"
  }

	fmt.Fprintf(w, "hello %v!\n", name)
}

func goodbye(w http.ResponseWriter, r *http.Request) {
  var name string

  if r.URL.Query().Has("name") {
    name = r.URL.Query().Get("name")
  } else {
    name = "friend"
  }

	fmt.Fprintf(w, "goodbye %v!\n", name)
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

    fmt.Fprintf(w, strconv.FormatFloat(difference.Abs().Hours()/24, 'f', 6, 64) + "\n")
  } else {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "exit\n")
  }
}

func brrrcalc(w http.ResponseWriter, r *http.Request) {
  if !r.URL.Query().Has("d") || !r.URL.Query().Has("s") {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Duration/Speed not given.\n")
  } else {
    duration, _ := strconv.ParseFloat(r.URL.Query().Get("d"), 32)
    speed, _ := strconv.ParseFloat(r.URL.Query().Get("s"), 32)

    result := strconv.FormatFloat(duration/speed, 'f', 1, 32)

    fmt.Fprintf(w, result + "\n")
  }
}

func main() {
  http.HandleFunc("/", root)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye", goodbye)
	http.HandleFunc("/datecalc", datecalc)
	http.HandleFunc("/brrrcalc", brrrcalc)

  fmt.Println("server started on :3000!")

	err := http.ListenAndServe(":3000", nil)
  if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed!\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}