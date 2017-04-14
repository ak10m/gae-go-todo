package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	Title string `json:"title"`
}

func init() {
	http.HandleFunc("/todo", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "[]")
	case "POST":
		task := Task{Title: r.FormValue("title")}
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
