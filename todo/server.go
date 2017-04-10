package todo

import (
  "fmt"
  "net/http"
  "google.golang.org/appengine"
  "google.golang.org/appengine/datastore"
)

func init() {
  http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello, TODO")
}
