package server

import (
  "os"
  "fmt"
  "net/http"
)

func init() {
  http.HandleFunc("/todo", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  //projectId := os.Getenv("GCP_PROJECT_ID")
  projectId := os.Getenv("USER")
  fmt.Fprintf(w, "[%s] %s %s", projectId, r.Method, r.URL.Path)
}
