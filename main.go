package main

import (
	"github.com/gin-gonic/gin"
)

type Tarefa struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Done  bool   `json:"done"`
}

var tarefas []Tarefa

func initTarefas() {
	tarefas = append(tarefas, Tarefa{ID: 1, Title: "Ir ao mercado", Body: "Comprar verduras na promoção de quarta-feira", Done: false})
	tarefas = append(tarefas, Tarefa{ID: 2, Title: "Fazer a tarefa da faculdade", Body: "terminar a tarefa de java", Done: false})
}

func main() {
	initTarefas()
	r := gin.Default()
	getRouter(r)
	r.Run(":8080")
}

func getRouter(r *gin.Engine) *gin.Engine {
	return r
}
