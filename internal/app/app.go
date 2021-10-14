package app

import (
	"fmt"
	"os"
	"time"

	"github.com/Chipazawra/czwrmailing/internal/services/auth"
	"github.com/Chipazawra/czwrmailing/internal/services/profile"
	"github.com/Chipazawra/czwrmailing/internal/todo"
	"github.com/Chipazawra/czwrmailing/pkg/config"
	"github.com/Chipazawra/czwrmailing/pkg/jwtmng"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func logfmt(params gin.LogFormatterParams) string {

	var statusColor, methodColor, resetColor string
	if params.IsOutputColor() {
		statusColor = params.StatusCodeColor()
		methodColor = params.MethodColor()
		resetColor = params.ResetColor()
	}

	if params.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		params.Latency = params.Latency - params.Latency%time.Second
	}

	return fmt.Sprintf("[CZWR-LOG] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		params.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, params.StatusCode, resetColor,
		params.Latency,
		params.ClientIP,
		methodColor, params.Method, resetColor,
		params.Path,
		params.ErrorMessage,
	)
}

func Run() {
	defer afterStart()

	if conf, err := config.LoadConf(); err == nil {
		// init engine
		r := gin.New()
		r.Use(gin.Recovery())
		// init logger
		if conf.Server.Log {

			if conf.Server.LogToFile {

				f, err := os.OpenFile(conf.Server.LogPath+"\\log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
				if err != nil {
					panic(err)
				}
				os.Stdout = f
				defer f.Close()
			}

			r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
				Formatter: logfmt,
				Output:    os.Stdout,
			}))
		}

		//init services
		tokenManager, err := jwtmng.NewManager(conf.Secret)

		if err != nil {
			panic(err)
		}

		authSvc := auth.NewAuth(tokenManager, &conf.AuthConf)
		authSvc.AddRoutes(r)

		profileSvc := profile.NewProfile(tokenManager)
		profileSvc.AddRoutes(r)

		todo := todo.NewToDO(nil)
		todo.AddRoutes(r)

		//init profiling
		pprof.Register(r, "debug/pprof")

		//run server
		err = r.Run(fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port))

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
