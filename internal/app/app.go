package app

import (
	"fmt"

	"github.com/Chipazawra/czwrmailing/internal/auth"
	"github.com/Chipazawra/czwrmailing/internal/config"
	"github.com/Chipazawra/czwrmailing/internal/jwtmng"
	"github.com/Chipazawra/czwrmailing/internal/profile"
	"github.com/Chipazawra/czwrmailing/internal/todo"
	"github.com/gin-contrib/pprof"
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

		authSvc := auth.NewAuth(tokenManager, &config.AuthConf)
		authSvc.AddRoutes(r)

		profileSvc := profile.NewProfile(tokenManager)
		profileSvc.AddRoutes(r)

		todo := todo.NewToDO(nil)
		todo.AddRoutes(r)

		pprof.Register(r, "dev/pprof")

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
