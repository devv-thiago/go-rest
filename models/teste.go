package models

type Tarefa struct {
	Titulo string `json:"titulo"`
	Descricao string `json:"descricao"`
}

var Tarefas []Tarefa