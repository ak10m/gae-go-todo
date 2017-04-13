package server

import (
  "fmt"
  "net/http"
)

func init() {
  http.HandleFunc("/todo", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello, Todo!! - %s %s", r.Method, r.URL.Path)
}
