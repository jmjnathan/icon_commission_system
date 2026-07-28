package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nathanchristiawan02/icon-commission-system-backend/config"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/middleware"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/router/auth"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/router/master"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Gagal load file .env")
	}

	config.ConnectDatabase()

	// Auth
	adminRepo := repository.NewAdminRepository(config.DB)
	authUsecase := usecase.NewAuthUsecase(adminRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	// Master Size
	masterSizeRepo := repository.NewMasterSizeRepository(config.DB)
	masterSizeUsecase := usecase.NewMasterSizeUsecase(masterSizeRepo)
	masterSizeHandler := handler.NewMasterSizeHandler(masterSizeUsecase)

	r := gin.Default()

	auth.RegisterAuthRoutes(r, authHandler)

	protected := r.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		auth.RegisterMeRoute(protected)
		master.RegisterMasterSizeRoutes(protected, masterSizeHandler)
	}

	r.Run(":8080")
}