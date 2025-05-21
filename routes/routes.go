package routes

import (
	"log"
	"net/http"

	"github.com/devv-thiago/go-rest.git/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/v1/tarefas", controllers.Tarefas)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
