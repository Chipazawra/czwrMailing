package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var whiteList = gin.Accounts{
	"user1": "love",
	"user2": "god",
	"user3": "sex",
}

func AddRoutes(r *gin.Engine) {
	authorized := r.Group("/", gin.BasicAuth(whiteList))
	authorized.GET("/login", BasicAuth)
	authorized.GET("/logout", Logout)
}

func BasicAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"secret": "The secret ingredient to the BBQ sauce is stiring it in an old whiskey barrel.",
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"secret": "logout.",
	})
}
