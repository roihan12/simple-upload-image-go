package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/controllers"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/helpers"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/middlewares"
	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/services"
)

type Controllers struct {
	UserController  controllers.UserController
	PhotoController controllers.PhotoController
}

type AuthMiddleware struct {
	AuthService helpers.JwtService
	UserService services.UserService
}

func NewRouter(c *Controllers, a *AuthMiddleware) *gin.Engine {
	router := gin.Default()
	api := router.Group("api/v1")

	userRoute := api.Group("/users")
	{
		userRoute.POST("/register", c.UserController.Register)
		userRoute.POST("/login", c.UserController.Login)
		userRoute.PUT("/:userId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.UserController.Update)
		userRoute.DELETE("/:userId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.UserController.Delete)
	}

	photoRoute := api.Group("/photos")
	{
		photoRoute.POST("/", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.Create)
		photoRoute.GET("/", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.GetPhoto)
		photoRoute.PUT("/:photoId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.Edit)
		photoRoute.DELETE("/:photoId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.Delete)
	}

	return router
}
