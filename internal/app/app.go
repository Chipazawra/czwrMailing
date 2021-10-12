package app

import (
	"fmt"

	"github.com/Chipazawra/czwrmailing/internal/auth"
	"github.com/Chipazawra/czwrmailing/internal/config"
	"github.com/Chipazawra/czwrmailing/internal/jwtmng"
	"github.com/Chipazawra/czwrmailing/internal/todo"
	"github.com/gin-gonic/gin"
)

func Run() {
	defer afterStart()

	if config, err := config.LoadConf(); err == nil {
		r := gin.Default()

		tokenManager, err := jwtmng.NewManager(config.Secret)

		if err != nil {
			panic(err)
		}

		authService := auth.NewAuth(tokenManager, &config.AuthConf)
		authService.AddRoutes(r)

		todo := todo.NewToDO(nil)
		todo.AddRoutes(r)

		err = r.Run(fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port))

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
