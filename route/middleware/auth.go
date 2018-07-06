package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		if err := CheckToken(c); err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("token error"))
		}
		c.Next()
	}
}

func CheckToken(c *gin.Context) error {

	return nil
}