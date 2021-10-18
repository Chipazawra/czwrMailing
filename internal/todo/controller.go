package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ToDo struct {
}

type ToDoConf struct {
}

func New(c *ToDoConf) *ToDo {
	return &ToDo{}
}

func (t *ToDo) Register(r *gin.Engine) {
	r.GET("/todo", t.Pong)
}

func (t *ToDo) Pong(c *gin.Context) {
	c.String(http.StatusOK, "todo")
}
