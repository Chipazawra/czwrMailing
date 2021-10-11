package app

import (
	"github.com/Chipazawra/czwrmailing/internal/auth"
	"github.com/Chipazawra/czwrmailing/internal/todo"
	"github.com/gin-gonic/gin"
)

func Run() {
	defer afterStart()
	r := gin.Default()
	auth.AddRoutes(r)
	todo.AddRoutes(r)
	if err := r.Run(); err != nil {
		panic(err)
	}
}

func afterStart() {
	//TODO
}
