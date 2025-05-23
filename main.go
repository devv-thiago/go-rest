package main

import (
	"github.com/devv-thiago/go-rest.git/models"
	"github.com/devv-thiago/go-rest.git/routes"
)

func main() {
	var tarefa1 = models.Tarefa{
		Id:        1,
		Titulo:    "Limpar casa",
		Descricao: "Passar pano em todos os quartos",
	}
	var tarefa2 = models.Tarefa{
		Id:        2,
		Titulo:    "Lavar a louça",
		Descricao: "Lavar e secar toda a louça",
	}

	models.Tarefas = []models.Tarefa{tarefa1, tarefa2}

	routes.HandleRequest()
}
