package api_handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pm-cloudify/multi-container-app/server/internal/models"
	"github.com/pm-cloudify/multi-container-app/server/internal/services"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GET /todos
func GetAllTodos(c *gin.Context) {
	todos, err := services.GetAllTodos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, &gin.H{
			"message": "failed to fetch todos.",
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// GET /todos/:id
func GetTodo(c *gin.Context) {
	id := c.Param("id")

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, &gin.H{
			"message": "data not found. invalid id.",
		})
		return
	}

	todo, err := services.GetTodo(&models.QueryTodo{
		ID: objectID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, &gin.H{
			"message": "failed to fetch todos.",
		})
		return
	}

	if todo == nil {
		c.JSON(http.StatusNotFound, &gin.H{
			"message": "data not found.",
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// POST /todos
func PostTodo(c *gin.Context) {
	data := models.NewTodoData{}

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, &gin.H{
			"message": "cannot process given data.",
		})
		return
	}

	if data.Title == "" {
		c.JSON(http.StatusUnprocessableEntity, &gin.H{
			"message": "title should not be empty.",
		})
		return
	}

	err := services.CreateNewTODO(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &gin.H{
			"message": "cannot created new data",
		})
		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"message": "new todo created successfully.",
	})
}

// TODO: complete bellow api handlers

// Put /todos/:id
func PutTodo(c *gin.Context) {

	id := c.Param("id")

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, &gin.H{
			"message": "data does not exist. invalid id.",
		})
		return
	}

	data := models.NewTodoData{}

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, &gin.H{
			"message": "cannot process given data.",
		})
		return
	}

	if data.Title == "" {
		c.JSON(http.StatusUnprocessableEntity, &gin.H{
			"message": "title should not be empty.",
		})
		return
	}

	err = services.UpdateOneTodo(&models.QueryTodo{ID: objectID}, &data)
	if err != nil {
		if err.Error() == "not modified" {
			c.JSON(http.StatusNotFound, &gin.H{
				"message": "data does not exist.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, &gin.H{
			"message": "cannot access data.",
		})
		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"message": "todo updated successfully.",
	})
}

// DELETE /todos/:id
func DeleteTodo(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, &gin.H{
		"message": "feature not implemented!",
	})
}
