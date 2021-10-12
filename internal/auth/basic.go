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
	authorized.GET("/login", a.BasicAuth)
	authorized.GET("/logout", a.Logout)
}

func (a *Auth) BasicAuth(c *gin.Context) {

	token, _ := a.TokenManager.NewJWT("usr", time.Duration(a.config.JwtTTL))
	refresh, _ := a.TokenManager.NewRefreshToken()

	c.SetCookie("access", token, 60, "/", "localhost", false, true)
	c.SetCookie("refresh", refresh, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"secret": "set cookie.",
	})
}

func (a *Auth) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"secret": "logout.",
	})
}
