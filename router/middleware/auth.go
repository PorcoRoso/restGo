package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/porcorosso/restGo/handler"
	"github.com/porcorosso/restGo/pkg/errno"
	"github.com/porcorosso/restGo/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
