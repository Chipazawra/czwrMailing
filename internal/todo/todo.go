package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	r.GET("/todo", Pong)
}

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
