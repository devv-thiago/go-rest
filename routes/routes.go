package routes

import (
	"log"
	"net/http"

	"github.com/devv-thiago/go-rest.git/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/v1/tarefas", controllers.Tarefas).Methods("Get")
	r.HandleFunc("/api/v1/tarefas/{id}", controllers.RetornaTarefa).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
