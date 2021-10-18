package app

import (
	"fmt"
	"os"
	"time"

	_ "github.com/Chipazawra/czwrmailing/doc"
	"github.com/Chipazawra/czwrmailing/internal/dbcontext/inmemoryctx"
	"github.com/Chipazawra/czwrmailing/internal/services/auth"
	"github.com/Chipazawra/czwrmailing/internal/services/profile"
	"github.com/Chipazawra/czwrmailing/internal/services/receivers"
	"github.com/Chipazawra/czwrmailing/internal/todo"
	"github.com/Chipazawra/czwrmailing/pkg/config"
	"github.com/Chipazawra/czwrmailing/pkg/jwtmng"
	"github.com/Chipazawra/czwrmailing/pkg/pprofwrapper"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

// @title czwrMailing service
// @version 1.0
// @description This is a sample mailing servivce.
func Run() {

	defer afterStart()

	if conf, err := config.Load(); err == nil {

		g := gin.New()
		g.Use(gin.Recovery())

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

			g.Use(gin.LoggerWithConfig(gin.LoggerConfig{
				Formatter: logfmt,
				Output:    os.Stdout,
			}))
		}

		tm, err := jwtmng.New(conf.Secret)

		if err != nil {
			panic(err)
		}

		//init db context
		db := inmemoryctx.New()

		//init services
		as := auth.New(tm, &conf.AuthConf)
		ps := profile.New(tm)
		ts := todo.New(nil)
		pw := pprofwrapper.New()
		rc := receivers.New(db)

		//register services
		as.Register(g)
		ps.Register(g)
		ts.Register(g)
		pw.Register(g)
		rc.Register(g)

		//swagger
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		//run server
		err = g.Run(fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port))

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
