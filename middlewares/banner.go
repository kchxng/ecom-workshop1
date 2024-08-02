package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewBanner() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("*********************************************")
		fmt.Println("*              auth-service                 *")
		fmt.Println("*         Welcome to go-gin App             *")
		fmt.Println("*********************************************")
		c.Next()
	}
}
