package main

import (
	"lld-examples/internal/taskmanager/handler"
	"lld-examples/internal/taskmanager/router"
	"lld-examples/internal/taskmanager/service"
	"log"
	"net/http"
)

func main() {

	taskService := service.NewTaskService()
	taskHandler := handler.NewTaskHandler(taskService)

	r := router.SetupRouter(taskHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
