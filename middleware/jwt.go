package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func AuthJWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// extract jwt, Authorization: "Bearer " + token
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" { // return in advance
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "Request header has no auth",
			})
			ctx.Abort()
			return
		}
		// parse jwt
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "Request header auth wrong format",
			})
			ctx.Abort()
			return
		}
		mc, err := ParseToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "Invalid token",
			})
			ctx.Abort()
			return
		}
		// if needed, inject username into context
		ctx.Set("username", mc.Username)
		// pass request
		ctx.Next()
	}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour

var MySecret = []byte("gin-jwt")

// generate token
func GenerateToken(username string) (string, error) {
	// create self claims
	c := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "project",
		},
	}
	// specific sign method to create obj
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return t.SignedString(MySecret)
}

// parse token
func ParseToken(tokenString string) (*MyClaims, error) {
	// parse token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return MySecret, nil
		},
	)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid token")
}
