package controllers

import (
	"my_first_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Tasks = []models.Task{
	{ID: 1, Title: "Aprender Go", Done: false},
	{ID: 2, Title: "Construir API", Done: true},
}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, Tasks)
}

func CreateTask(c *gin.Context) {
	var new models.Task

	if err := c.ShouldBindJSON(&new); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	new.ID = len(Tasks) + 1
	Tasks = append(Tasks, new)

	c.JSON(http.StatusCreated, new)
}

func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")

	var index = -1

	for i, t := range Tasks {

		if idParam == strconv.Itoa(t.ID) {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task doesn't exist"})
		return
	}

	Tasks = append(Tasks[:index], Tasks[index+1:]...)
	c.JSON(http.StatusOK, gin.H{"mssg": "Tasks was deleted"})
}
