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

	// === Auth ===
	authRoutes := api.Group("/auth")
	authRoutes.POST("/register", authHandler.Register)
	authRoutes.POST("/login", authHandler.Login)

	// === Protected Routes ===
	protectedRoutes := api.Group("")
	protectedRoutes.Use(middleware.AuthMidleware())
	protectedRoutes.POST("/auth/logout", authHandler.Logout)

	// === User ===
	userRoutes := api.Group("/users")
	userRoutes.GET("/profile", userHandler.GetProfile)
}
