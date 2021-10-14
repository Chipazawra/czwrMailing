package profile

import (
	"net/http"

	"github.com/Chipazawra/czwrmailing/internal/jwtmng"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Profile struct {
	TokenManager *jwtmng.Mng
}

func NewProfile(tm *jwtmng.Mng) *Profile {
	return &Profile{TokenManager: tm}
}

func (p *Profile) AddRoutes(r *gin.Engine) {
	g := r.Group("/profile")
	g.GET("/i", p.iHandler)
	g.GET("/me", p.meHandler)
}

func (p *Profile) iHandler(c *gin.Context) {
	p.pHandler(c, p.TokenManager.ParseToken)
}

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
