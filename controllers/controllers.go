package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/devv-thiago/go-rest.git/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func Tarefas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Tarefas)
}

func RetornaTarefa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, tarefa := range models.Tarefas {
		if strconv.Itoa(tarefa.Id) == id {
			json.NewEncoder(w).Encode(tarefa)
		}
	}
}
