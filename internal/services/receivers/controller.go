package receivers

import (
	"github.com/gin-gonic/gin"
)

func (r *Receivers) Register(g *gin.Engine) {
	g.POST("/reciviers/:usr/*receiver", r.CreateHandler)
	g.GET("/reciviers/:usr", r.ReadHandler)
	g.PATCH("/reciviers/:usr/*id", r.UpdateHandler)
	g.DELETE("/reciviers/:usr/*id", r.DeleteHandler)
}

func (r *Receivers) CreateHandler(c *gin.Context) {

}

func (r *Receivers) ReadHandler(c *gin.Context) {

}

func (r *Receivers) UpdateHandler(c *gin.Context) {

}

func (r *Receivers) DeleteHandler(c *gin.Context) {

}
