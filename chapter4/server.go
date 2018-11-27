package main

import (
  "fmt"
  "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "hello world")
}

func main() {
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":3000", nil)
}

