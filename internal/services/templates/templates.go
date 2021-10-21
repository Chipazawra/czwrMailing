package templates

import (
	"html/template"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type Templates struct {
}

func New() *Templates {
	return &Templates{}
}

func (t *Templates) Register(g *gin.Engine) {
	g.POST("/upload_template", t.uploadTemplateHandler)
}

func (t *Templates) uploadTemplateHandler(c *gin.Context) {

	raw, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	tmpl := template.New("tmpl")

	_, err = tmpl.Parse(string(raw))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	rx, _ := regexp.Compile(`{{ \.(.*?)}}`)
	as := rx.FindAllStringSubmatch(string(raw), -1)
	c.JSON(http.StatusOK, as)

}
