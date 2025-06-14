package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.comluizhpferreirago-standards/model"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(model.GetAllTasks())
	case http.MethodPost:
		body, _ := io.ReadAll(r.Body)
		var task model.Task
		json.Unmarshal(body, &task)
		model.AddTask(&task)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func TaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		task, found := model.GetTaskByID(id)
		if !found {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(task)
	case http.MethodPut:
		body, _ := io.ReadAll(r.Body)
		var updatedTask model.Task
		json.Unmarshal(body, &updatedTask)
		updated := model.UpdateTask(id, updatedTask)
		if !updated {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(updatedTask)
	case http.MethodDelete:
		if !model.DeleteTask(id) {
			http.NotFound(w, r)
			return
		}
		fmt.Fprint(w, "Tarefa deletada")
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}
