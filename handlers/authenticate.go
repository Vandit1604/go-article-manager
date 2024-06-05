package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowSignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// for testing purpose using password store as DB
var DB = map[string]string{}

func AuthenticateSignUp(c *gin.Context) {
	var Username string
	var Password string

	Username = c.PostForm("username")
	Password = c.PostForm("password")

	DB[Username] = Password

	c.HTML(http.StatusOK, "validate-response.html", gin.H{
		"greet":    "Welcome",
		"response": "Aboard!",
		"message":  "Redirecting you to blogs",
	})
}

func AuthenticateLogin(c *gin.Context) {
	Password := c.PostForm("password")

	for key := range DB {
		if DB[key] == Password {
			c.HTML(http.StatusTemporaryRedirect, "validate-response.html", gin.H{
				"greet":    "Welcome",
				"response": "Back!",
			})
			c.Redirect(301, "/")
		}
	}
}
