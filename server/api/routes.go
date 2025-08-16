package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api_handlers "github.com/pm-cloudify/multi-container-app/server/api/handlers"
)

func AttachHandler(router *gin.Engine) {

	// todos handler:
	{ // get
		router.Handle(http.MethodGet, "/todos", api_handlers.GetAllTodos)
		router.Handle(http.MethodGet, "/todos/:id", api_handlers.GetTodo)
		// post
		router.Handle(http.MethodPost, "/todos", api_handlers.PostTodo)
		// put
		router.Handle(http.MethodPost, "/todos/:id", api_handlers.PutTodo)
		// delete
		router.Handle(http.MethodDelete, "/todos/:id", api_handlers.DeleteTodo)
	}

}
