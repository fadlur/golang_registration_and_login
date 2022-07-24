package routes

import (
	"registration/handler"
	"registration/middleware"
	"registration/repository"
	"registration/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApiRoute(db *gorm.DB) *gin.Engine {
	roleRepository := repository.InitRoleRepository(db)
	roleService := services.InitRoleService(roleRepository)
	roleHandler := handler.InitRoleHandler(roleService)

	userRepository := repository.InitUserRepository(db)
	userService := services.InitUserService(userRepository)
	userHandler := handler.Inituserhandler(userService)
	
	router := gin.Default()
	api := router.Group("/api")
	{
		role := api.Group("/role")
		{
			role.GET("/", roleHandler.GetRoles)
			role.GET("/:id", roleHandler.GetRoleById)
			role.PATCH("/:id", roleHandler.UpdateRole)
			role.DELETE("/:id", roleHandler.DeleteRole)
			role.POST("/", roleHandler.CreateRole)
		}
		user := api.Group("/user")
		{
			user.GET("/:id", userHandler.Register)
			user.POST("/", userHandler.Register)
			user.POST("/login", userHandler.Login)
		}
		api.GET("/dashboard", middleware.Auth(), handler.Dashboard)
	}
	return router
}