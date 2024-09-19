package main

import (
	"net/http"
	"strconv"

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
	rotasTarefa := r.Group("/tarefas")
	rotasTarefa.GET("", getTarefas)
	rotasTarefa.GET("/:id", getTarefaById)
	return r
}

func getTarefas(c *gin.Context) {
	c.JSON(http.StatusOK, tarefas)
}

func getTarefaById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	for _, tarefa := range tarefas {
		if tarefa.ID == id {
			c.JSON(http.StatusOK, tarefa)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
}
