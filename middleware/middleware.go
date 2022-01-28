package middleware

import (
	"Final-ProjectBDS-Sanbercode-Golang-Batch-31/utils/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userMiddleware struct {
	authService auth.Service
}

func NewUserMiddleware(authService auth.Service) *userMiddleware {
	return &userMiddleware{authService}
}

func (u userMiddleware) JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := u.authService.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		Id, err1 := u.authService.ExtractTokenID(c)
		if err1 != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		c.Set("currentUser", Id)
		c.Next()
	}
}
