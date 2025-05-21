package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devv-thiago/go-rest.git/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func Tarefas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Tarefas)
}
