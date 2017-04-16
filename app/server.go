package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
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
	http.HandleFunc("/todo", ResourcesHandler)
	http.HandleFunc("/todo/:id", ResourceHandler)
}

func ResourcesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		tasks := GetTasks()
		out, err := json.Marshal(tasks)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(out))
	case http.MethodPost:
		task := Task{
			Title:     r.FormValue("title"),
			CreatedAt: time.Now(),
		}
		_, err := createTask(ctx, task)

		if err != nil {
			ErrorResponse(w, 422)
			return
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
		ErrorResponse(w, http.StatusMethodNotAllowed)
	}
}

func ResourceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	taskId := "anyTaskId"
	task := GetTask(taskId)

	// TODO handle not_found
	if false {
		ErrorResponse(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		out, err := json.Marshal(task)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(out))
	case http.MethodPut:
		// TODO update task
		out, err := json.Marshal(task)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, string(out))
	case http.MethodDelete:
		// TODO delete task
		w.WriteHeader(http.StatusNoContent)
	default:
		ErrorResponse(w, http.StatusMethodNotAllowed)
	}
}

func ErrorResponse(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	e := ErrorJson{Message: http.StatusText(statusCode)}
	out, _ := json.Marshal(e)
	fmt.Fprintln(w, string(out))
}

func GetTasks() []Task {
	tasks := []Task{
		{Title: "task 1", CreatedAt: time.Now()},
		{Title: "task 2", CreatedAt: time.Now()},
	}
	return tasks
}

func GetTask(id string) Task {
	task := Task{
		Title:     "task",
		CreatedAt: time.Now(),
	}
	return task
}

func createTask(ctx context.Context, task Task) (string, error) {
	if false {
		return "", errors.New("Validation Error")
	}
	return "newId", nil
}
