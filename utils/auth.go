package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("wiauwonosobo")

func getRole(username string) string {
	if username == "ben1" {
		return "ben1"
	}
	return "employee"
}

// coba testing create token dgn panggil API test
func CreateToken(username string) (string, error) {
	// jwt.SigningMethodHS256 ==> apa bedanya kapan pakenya ?
	// jwt.SigningMethodES256 ==> apa bedanya kapan pakenya ?

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// nanti add another info
		"issuer":   username,
		"expired":  time.Now().Add(time.Hour).Unix(),
		"issuedAt": time.Now().Unix(),
		"audience": getRole(username),
	})
	fmt.Printf("claim jwt: %+v", claims)

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	// check verification error
	if err != nil {
		return nil, err
	}
	// check if token valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	// return token
	return token, nil

}

func AuthenticateMiddleware(c *gin.Context) {
	// Retrieve the token from the cookie
	tokenString, err := c.Cookie("token")
	if err != nil {
		fmt.Println("Token missing in cookie")

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Verify the token
	token, err := VerifyToken(tokenString)
	if err != nil {
		fmt.Printf("Token verification failed: %v\\n", err)
		c.AbortWithStatusJSON(http.StatusNonAuthoritativeInfo, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Print information about the verified token
	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)

	// Continue with the next middleware or route handler
	c.Next()
}
