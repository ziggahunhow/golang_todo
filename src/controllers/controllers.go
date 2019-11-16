package controllers

import (
	models "Todo/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todo []models.Todo
	err := models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	if todo.Description == "" || todo.Title == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	err := models.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
