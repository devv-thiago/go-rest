package models

type Tarefa struct {
	Id int `json:"id"`
	Titulo string `json:"titulo"`
	Descricao string `json:"descricao"`
}

var Tarefas []Tarefa