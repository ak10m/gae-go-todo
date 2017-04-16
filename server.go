package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

type ErrorJson struct {
	Message string `json:"message"`
}

func init() {
	// TODO routing
	http.HandleFunc("/todo", resourcesHandler)
	http.HandleFunc("/todo/:id", resourceHandler)
}

func resourcesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		tasks := getTasks()
		out, err := json.Marshal(tasks)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(out))
	case "POST":
		task := &Task{
			Title:     r.FormValue("title"),
			CreatedAt: time.Now(),
		}
		out, err := json.Marshal(task)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// TODO create task
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, string(out))
	default:
		errorResponse(w, http.StatusMethodNotAllowed)
	}
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taskId := "anyTaskId"
	task := getTask(taskId)

	// TODO handle not_found
	if false {
		errorResponse(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		out, err := json.Marshal(task)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(out))
	case "PUT":
		// TODO update task
		out, err := json.Marshal(task)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, string(out))
	case "DELETE":
		// TODO delete task
		w.WriteHeader(http.StatusNoContent)
	default:
		errorResponse(w, http.StatusMethodNotAllowed)
	}
}

func errorResponse(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	e := ErrorJson{Message: http.StatusText(statusCode)}
	out, _ := json.Marshal(e)
	fmt.Fprintln(w, string(out))
}

func getTasks() []Task {
	tasks := []Task{
		{Title: "task 1", CreatedAt: time.Now()},
		{Title: "task 2", CreatedAt: time.Now()},
	}
	return tasks
}

func getTask(id string) Task {
	task := Task{
		Title:     "task",
		CreatedAt: time.Now(),
	}
	return task
}

func newTask(title string) Task {
	task := Task{
		Title:     title,
		CreatedAt: time.Now(),
	}
	return task
}
