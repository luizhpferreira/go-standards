package main

import (
	"log"
	"net/http"

	"github.comluizhpferreirago-standards/handler"
)

func main() {
	http.HandleFunc("/tasks", handler.TasksHandler) //GET, POST
	http.HandleFunc("/tasks/", handler.TaskByIDHandler)

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
