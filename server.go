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
	// TODO routing
	http.HandleFunc("/todo", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		tasks := []Task{
			{ Title: "task 1" },
			{ Title: "task 2" },
		}
		out, err := json.Marshal(tasks)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		fmt.Fprintln(w, string(out))
	case "POST":
		task := Task{Title: r.FormValue("title")}
		out, err := json.Marshal(task)
		if err != nil {
			w.WriteHeader(422)
			return
		}
		w.WriteHeader(201)
		fmt.Fprintln(w, string(out))
	case "DELETE":
		// TODO delete task
		w.WriteHeader(204)
	default:
		w.WriteHeader(400)
	}
}
