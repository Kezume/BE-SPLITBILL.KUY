package router

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/handler"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/middleware"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/repository"
	"github.com/Kezume/BE-SPLITBILL.KUY/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	// === dependency injection ===
	userRepo := repository.NewUserRepository()
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	dashboardRepo := repository.NewDashboardRepository()
	dashboardService := service.NewDashboardService(dashboardRepo)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)

	groupRepo := repository.NewGroupRepository()
	groupService := service.NewGroupService(groupRepo)
	groupHandler := handler.NewGroupHandler(groupService)

	// === Auth ===
	authRoutes := api.Group("/auth")
	authRoutes.POST("/register", authHandler.Register)
	authRoutes.POST("/login", authHandler.Login)

	// === Protected Routes ===
	protectedRoutes := api.Group("")
	protectedRoutes.Use(middleware.AuthMidleware())
	protectedRoutes.POST("/auth/logout", authHandler.Logout)

	// === User (Protected) ===
	userRoutes := protectedRoutes.Group("/users")
	userRoutes.GET("/profile", userHandler.GetProfile)
	userRoutes.PUT("/profile", userHandler.UpdateProfile)
	userRoutes.DELETE("/profile", userHandler.DeleteProfile)

	// === Dashboard (Protected) ===
	dashboardRoutes := protectedRoutes.Group("/dashboard")
	dashboardRoutes.GET("/", dashboardHandler.GetDashboard)

	// === Group (Protected) ===
	groupRoutes := protectedRoutes.Group("/groups")
	groupRoutes.POST("/", groupHandler.CreateGroup)
	groupRoutes.GET("/list", groupHandler.GetListGroup)
	groupRoutes.GET("/:id", groupHandler.GetGroupDetail)
	groupRoutes.DELETE("/:id", groupHandler.DeleteGroup)
	groupRoutes.POST("/join", groupHandler.JoinGroup)
}
