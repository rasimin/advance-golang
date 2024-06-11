package router

import (
	"advance/handler"
	"advance/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// usersPublicEndpoint := r.Group("/users")
	// usersPublicEndpoint.GET("/:id", handler.GetUser)
	// usersPublicEndpoint.GET("/", handler.GetAllUsers)

	// usersPrivateEndpoint := r.Group("/users")
	// usersPrivateEndpoint.Use(middleware.AuthMiddleware())
	// usersPrivateEndpoint.POST("/", handler.CreateUser)
	// usersPrivateEndpoint.PUT("/:id", handler.UpdateUser)
	// usersPrivateEndpoint.DELETE("/:id", handler.DeleteUser)

	r.GET("/", handler.RootHandler)

	// Tambahkan middleware AuthMiddleware ke rute yang memerlukan autentikasi
	privateEndpoint := r.Group("/private")
	privateEndpoint.Use(middleware.AuthMiddleware())
	{
		privateEndpoint.POST("/post", handler.PostHandler)
	}
}
