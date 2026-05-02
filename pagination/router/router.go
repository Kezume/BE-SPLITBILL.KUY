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

	// === User / Auth ===
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)

	authRoutes := api.Group("/auth")
	authRoutes.POST("/register", authHandler.Register)
	authRoutes.POST("/login", authHandler.Login)

	// === Protected Routes ===
	protectedRoutes := api.Group("")
	protectedRoutes.Use(middleware.AuthMidleware())
	protectedRoutes.POST("/auth/logout", authHandler.Logout)
}
