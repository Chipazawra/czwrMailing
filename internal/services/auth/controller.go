package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *Auth) Register(r *gin.Engine) {
	authorized := r.Group("/", gin.BasicAuth(a.config.Users))
	authorized.GET("/login", a.loginHandler)
	r.GET("/logout", a.logoutHandler)
}

// login godoc
// @Summary login in service
// @Tags auth
// @Description get auth data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /login [get]
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

// logout godoc
// @Summary logout from service
// @Tags auth
// @Description clear auth data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /logout [get]
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
