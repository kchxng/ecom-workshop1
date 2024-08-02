package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	}
}
