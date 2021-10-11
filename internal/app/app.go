package app

import (
	"fmt"

	"github.com/Chipazawra/czwrmailing/internal/config"

	"github.com/Chipazawra/czwrmailing/internal/auth"
	"github.com/Chipazawra/czwrmailing/internal/todo"
	"github.com/gin-gonic/gin"
)

func Run() {
	defer afterStart()

	if config, err := config.LoadConf(); err == nil {

		r := gin.Default()
		auth.AddRoutes(r, gin.Accounts(config.Auth.WhiteList))
		todo.AddRoutes(r)

		err := r.Run(fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port))

		if err != nil {
			panic(err)
		}

	} else {
		panic(err)
	}
}

func afterStart() {
	//TODO
}
