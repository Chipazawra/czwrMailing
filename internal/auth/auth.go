package auth

import (
	"net/http"
	"time"

	"github.com/Chipazawra/czwrmailing/internal/jwtmng"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	TokenManager *jwtmng.Mng
	config       AuthConf
}

type AuthConf struct {
	Users      map[string]string `yaml:"users"`
	JwtTTL     int               `yaml:"jwtttl"`
	RefreshTTL int               `yaml:"refreshttl"`
	Secret     string            `yaml:"secret"`
}

func NewAuth(tm *jwtmng.Mng, c *AuthConf) *Auth {
	return &Auth{TokenManager: tm, config: *c}
}

func (a *Auth) AddRoutes(r *gin.Engine) {
	authorized := r.Group("/", gin.BasicAuth(a.config.Users))
	authorized.GET("/login", a.loginHandler)
	r.GET("/logout", a.logoutHandler)
}

func (a *Auth) loginHandler(c *gin.Context) {

	token, _ := a.TokenManager.NewJWT("usr", time.Duration(a.config.JwtTTL))
	refresh, _ := a.TokenManager.NewRefreshToken()

	c.SetCookie("access", token, a.config.JwtTTL, "/", "localhost", false, true)
	c.SetCookie("refresh", refresh, a.config.RefreshTTL, "/", "localhost", false, true)

	rurl, rdct := c.GetQuery("redirect_uri")

	if rdct {
		c.Redirect(http.StatusPermanentRedirect, rurl)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "login.",
		})
	}

}

func (a *Auth) logoutHandler(c *gin.Context) {

	c.SetCookie("access", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh", "", -1, "/", "localhost", false, true)

	rurl, rdct := c.GetQuery("redirect_uri")

	if rdct {
		c.Redirect(http.StatusPermanentRedirect, rurl)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "logout.",
		})
	}
}
