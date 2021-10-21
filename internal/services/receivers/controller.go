package receivers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Receivers) Register(g *gin.Engine) {
	g.POST("/reciviers/:usr/:receiver", r.CreateHandler)
	g.GET("/reciviers/:usr", r.ReadHandler)
	g.PATCH("/reciviers/:usr/:id/:receiver", r.UpdateHandler)
	g.DELETE("/reciviers/:usr/:id", r.DeleteHandler)
}

func (r *Receivers) CreateHandler(c *gin.Context) {

	usr := c.Param("usr")
	receiver := c.Param("receiver")
	id, err := r.Create(usr, receiver)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (r *Receivers) ReadHandler(c *gin.Context) {

	usr := c.Param("usr")

	receivers, err := r.Read(usr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"receivers": receivers,
	})
}

func (r *Receivers) UpdateHandler(c *gin.Context) {

	usr := c.Param("usr")
	receiver := c.Param("receiver")
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	err = r.Update(usr, uint(id), receiver)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}

func (r *Receivers) DeleteHandler(c *gin.Context) {

	usr := c.Param("usr")

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	err = r.Delete(usr, uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
}
