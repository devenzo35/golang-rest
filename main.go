package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var Tasks = []Task{
	{ID: 1, Title: "Aprender Go", Done: false},
	{ID: 2, Title: "Construir API", Done: true},
}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, Tasks)
}

func main() {
	r := gin.Default()

	r.GET("/tasks", getTasks)

	r.Run()
}
