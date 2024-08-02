package router

import (
	app_user "github.com/kchxng/ecom-api/api/users"
	"github.com/kchxng/ecom-api/middlewares"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	if os.Getenv("NDOE_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(middlewares.NewCorsConfig())
	router.Use(middlewares.NewRequestID())
	p := middlewares.NewCustomPrometheus("gin")
	p.Use(router)
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Hello World",
		})
	})

	// router.GET("/users", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": http.StatusOK,
	// 		"data":   "USER",
	// 	})
	// })

	// if os.Getenv("NDOE_ENV") == "development" {
	// 	userRepository := app_user.NewCustomRepository(db)
	// 	userService := app_user.NewCustomService(userRepository)
	// 	app_user.NewCustomHandler(router.Group("/users"), userService)
	// }

	userRepository := app_user.NewUserRepository(db)
	userService := app_user.NewUserService(userRepository)
	app_user.NewUserHandler(router.Group("/users"), userService)

	return router
}
