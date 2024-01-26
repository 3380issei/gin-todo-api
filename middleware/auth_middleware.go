package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware interface {
	JWTAuth(c *gin.Context)
	CORS() gin.HandlerFunc
}

type authMiddleware struct{}

func NewAuthMiddleware() AuthMiddleware {
	return &authMiddleware{}
}

func (am *authMiddleware) JWTAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error extracting claims")
		return
	}

	userID, ok := claims["user_id"].(float64)

	if !ok {
		fmt.Println("Error extracting user_id")
		return
	}
	c.Set("user_id", uint(userID))
	c.Next()
}

func (am *authMiddleware) CORS() gin.HandlerFunc {
	return cors.New(cors.Config{

		AllowOrigins: []string{
			"http://localhost:3000",
		},

		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},

		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})
}
