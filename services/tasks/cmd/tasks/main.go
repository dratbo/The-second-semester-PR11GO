package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"example.com/pz11-graphql/internal/httpapi"
	"example.com/pz11-graphql/internal/service"
	"example.com/pz11-graphql/internal/task"
)

func main() {
	port := os.Getenv("TASKS_PORT")
	if port == "" {
		port = "8082"
	}

	repo := task.NewRepo()
	taskService := service.NewTaskService(repo)
	handler := httpapi.NewHandler(taskService)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"service": "tasks",
		})
	})

	mux.HandleFunc("/v1/tasks", handler.GetTasks)
	mux.HandleFunc("/v1/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetTaskByID(w, r)
		case http.MethodPatch:
			handler.PatchTask(w, r)
		case http.MethodDelete:
			handler.DeleteTask(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	addr := ":" + port
	log.Println("tasks service started on", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
