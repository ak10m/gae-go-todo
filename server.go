package server

import (
  "fmt"
  "net/http"
  "encoding/json"
)

type Task struct {
  Title string
}

func init() {
  http.HandleFunc("/todo", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  switch r.Method {
    case "GET":
      fmt.Fprintln(w, "fetch todo list")
    case "POST":
      task := Task{ Title: r.FormValue("title") }
      out, err := json.Marshal(task)
      if err != nil {
        w.WriteHeader(422)
        return
      }
      w.WriteHeader(201)
      fmt.Fprintln(w, string(out))
    default:
      w.WriteHeader(400)
  }
}
