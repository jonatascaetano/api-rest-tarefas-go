package main

import (
	"math/rand"
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
	rotasTarefa.POST("", postTarefa)
	rotasTarefa.DELETE("/:id", deleteTarefa)
	rotasTarefa.PUT("/:id", putTarefa)
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

func postTarefa(c *gin.Context) {
	var tarefa Tarefa
	c.BindJSON(&tarefa)

	tarefa.ID = generateUniqueID()
	tarefas = append(tarefas, tarefa)
	c.JSON(http.StatusCreated, tarefa)
}

func generateUniqueID() int {
	var ids []int
	for _, t := range tarefas {
		ids = append(ids, t.ID)
	}

	for {
		numeroAleatorio := rand.Intn(1000)

		found := false
		for _, id := range ids {
			if id == numeroAleatorio {
				found = true
				break
			}
		}

		if !found {
			return numeroAleatorio
		}
	}

}

func deleteTarefa(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	for index, tarefa := range tarefas {
		if tarefa.ID == id {
			tarefas = append(tarefas[:index], tarefas[index+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}
func putTarefa(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var tarefa Tarefa
	c.BindJSON(&tarefa)

	tarefaEncontrada := false
	for index, tarefaAtual := range tarefas {
		if tarefaAtual.ID == id {
			tarefas[index].Title = tarefa.Title
			tarefas[index].Body = tarefa.Body
			tarefas[index].Done = tarefa.Done
			tarefaEncontrada = true
			tarefa = tarefas[index]
			break
		}
	}

	if !tarefaEncontrada {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, tarefa)
}
