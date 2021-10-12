package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ToDo struct {
}

type ToDoConf struct {
}

func NewToDO(c *ToDoConf) *ToDo {
	return &ToDo{}
}

func (t *ToDo) AddRoutes(r *gin.Engine) {
	r.GET("/todo", t.Pong)
}

func (t *ToDo) Pong(c *gin.Context) {
	c.String(http.StatusOK, "todo")
}
