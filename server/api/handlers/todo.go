package api_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: complete bellow api handlers

// GET /todos
func GetAllTodos(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, &gin.H{
		"message": "feature not implemented!",
	})
}

// GET /todos/:id
func GetTodo(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, &gin.H{
		"message": "feature not implemented!",
	})
}

// POST /todos
func PostTodo(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, &gin.H{
		"message": "feature not implemented!",
	})
}

// Put /todos/:id
func PutTodo(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, &gin.H{
		"message": "feature not implemented!",
	})
}

// DELETE /todos/:id
func DeleteTodo(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, &gin.H{
		"message": "feature not implemented!",
	})
}
