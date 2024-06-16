package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(`secret`)

type Credentials struct {
	Username string
	Password string
}

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

func ShowSignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// for testing purpose using password store as DB
var DB = map[string]string{"username": "password"}

func AuthenticateSignUp(c *gin.Context) {
	var userCredentials Credentials

	userCredentials.Username = c.PostForm("username")
	userCredentials.Password = c.PostForm("password")

	DB[userCredentials.Username] = userCredentials.Password

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: userCredentials.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userCredentials.Username,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// this gives unsigned jwt token with claims and algo name
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// this signs the jwt token returning complete and signed jwt token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	c.HTML(http.StatusOK, "validate-response.html", gin.H{
		"greet":    "Welcome",
		"response": "Aboard!",
		"message":  "Redirecting you to blogs",
	})
}

func AuthenticateLogin(c *gin.Context) {
	var userCredentials Credentials

	userCredentials.Password = c.PostForm("password")
	userCredentials.Username = c.PostForm("username")

	if DB[userCredentials.Username] == userCredentials.Password {
		c.HTML(http.StatusTemporaryRedirect, "validate-response.html", gin.H{
			"greet":    "Welcome",
			"response": "Back!",
			"message":  "Redirecting you to blogs",
		})
	} else {
		c.HTML(http.StatusUnauthorized, "wrong-validate.html", gin.H{
			"greet":    "Wrong",
			"response": "Password",
		})
	}
}

func CheckUserAuthenticated(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.Redirect(http.StatusTemporaryRedirect, "/signup")
			c.Abort()
			return
		}
		// For any other type of error, return a bad request status
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	tknString := cookie.Value
	claims := &Claims{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknString, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.Redirect(http.StatusTemporaryRedirect, "/signup")
			c.Abort()
			return
		}

		c.Writer.WriteHeader(http.StatusBadRequest)
		return

	}
	if !tkn.Valid {
		c.Redirect(http.StatusTemporaryRedirect, "/signup")
		c.Abort()
		return
	}
}
