package profile

import (
	"net/http"

	"github.com/Chipazawra/czwrmailing/pkg/jwtmng"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Profile struct {
	TokenManager *jwtmng.Mng
}

func New(tm *jwtmng.Mng) *Profile {
	return &Profile{TokenManager: tm}
}

func (p *Profile) Register(g *gin.Engine) {
	gr := g.Group("/profile")
	gr.GET("/i", p.iHandler)
	gr.GET("/me", p.meHandler)
}

// profile i godoc
// @Summary show user info
// @Tags profile
// @Description get auth data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /profile/i [get]
func (p *Profile) iHandler(c *gin.Context) {
	p.pHandler(c, p.TokenManager.ParseToken)
}

// profile me godoc
// @Summary valid jwt token and show user info
// @Tags profile
// @Description get auth data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /profile/me [get]
func (p *Profile) meHandler(c *gin.Context) {
	p.pHandler(c, p.TokenManager.ValidToken)
}

func (p *Profile) pHandler(c *gin.Context, fn func(val string) (jwt.Claims, error)) {

	ac, err := c.Request.Cookie("access")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "access token not found.",
		})
		return
	}

	data, err := fn(ac.Value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "invalid token",
		})
		return
	} else {
		c.JSON(http.StatusOK, data)
	}

}
