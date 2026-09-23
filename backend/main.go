package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nathanchristiawan02/icon-commission-system-backend/config"

	adminHandler "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/admin"
	clientHandler "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/client"
	masterHandler "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/master"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/router/auth"
	transaction "github.com/nathanchristiawan02/icon-commission-system-backend/internal/router/client"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/router/commission"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/router/master"

	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/middleware"

	adminRepo "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/admin"
	clientRepo "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/client"
	masterRepo "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/master"

	adminUsecase "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/admin"
	clientUsecase "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/client"
	masterUsecase "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/master"

	commissionHandlerPkg "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/commission"
	uploadHandler "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/upload"
	commissionRepoPkg "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/commission"
	commissionUsecasePkg "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/commission"

	cashoutHandlerPkg "github.com/nathanchristiawan02/icon-commission-system-backend/internal/handler/cashflow/cashout"
	cashoutRepoPkg "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/cashflow/cashout"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/router/cashflow/cashout"
	cashoutUsecasePkg "github.com/nathanchristiawan02/icon-commission-system-backend/internal/usecase/cashflow/cashout"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Gagal load file .env")
	}

	config.ConnectDatabase()

	// Auth
	adminRepository := adminRepo.NewAdminRepository(config.DB)
	authUC := adminUsecase.NewAuthUsecase(adminRepository)
	authH := adminHandler.NewAuthHandler(authUC)

	// Master Size
	masterSizeRepository := masterRepo.NewMasterSizeRepository(config.DB)
	masterSizeUseCase := masterUsecase.NewMasterSizeUsecase(masterSizeRepository)
	masterSizeH := masterHandler.NewMasterSizeHandler(masterSizeUseCase)

	// Master Material
	masterMaterialRepository := masterRepo.NewMasterMaterialRepository(config.DB)
	masterMaterialUC := masterUsecase.NewMasterMaterialUsecase(masterMaterialRepository)
	masterMaterialH := masterHandler.NewMasterMaterialHandler(masterMaterialUC)

	// Master Saints
	masterSaintsRepository := masterRepo.NewMasterSaintsRepository(config.DB)
	masterSaintsUC := masterUsecase.NewMasterSaintsUsecase(masterSaintsRepository)
	masterSaintsH := masterHandler.NewMasterSaintsHandler(masterSaintsUC)

	// Master Style
	masterStyleRepository := masterRepo.NewMasterStyleRepository(config.DB)
	masterStyleUC := masterUsecase.NewMasterStyleUsecase(masterStyleRepository)
	masterStyleH := masterHandler.NewMasterStyleHandler(masterStyleUC)

	// Material Component
	materialComponentRepository :=
		masterRepo.NewMaterialComponentRepository(config.DB)

	materialComponentUseCase :=
		masterUsecase.NewMaterialComponentUsecase(
			materialComponentRepository,
		)

	materialComponentH :=
		masterHandler.NewMaterialComponentHandler(
			materialComponentUseCase,
		)

	materialComponentVariantRepository :=
		masterRepo.NewMaterialComponentVariantRepository(config.DB)

	materialComponentVariantUC :=
		masterUsecase.NewMaterialComponentVariantUsecase(
			materialComponentVariantRepository,
		)

	materialComponentVariantH :=
		masterHandler.NewMaterialComponentVariantHandler(
			materialComponentVariantUC,
		)

	// Client
	clientRepository := clientRepo.NewClientRepository(config.DB)
	clientUC := clientUsecase.NewClientUsecase(clientRepository)
	clientH := clientHandler.NewClientHandler(clientUC)

	// Commission
	commissionRepository := commissionRepoPkg.NewCommissionRepository(config.DB)
	commissionUC := commissionUsecasePkg.NewCommissionUsecase(commissionRepository)
	commissionH := commissionHandlerPkg.NewCommissionHandler(commissionUC)

	commissionItemMaterialRepository :=
		commissionRepoPkg.NewCommissionItemMaterialRepository(config.DB)

	commissionItemMaterialUseCase :=
		commissionUsecasePkg.NewCommissionItemMaterialUsecase(
			commissionItemMaterialRepository,
			masterRepo.NewMaterialComponentVariantRepository(config.DB),
		)

	commissionItemMaterialHandler :=
		commissionHandlerPkg.NewCommissionItemMaterialHandler(
			commissionItemMaterialUseCase,
		)

	// Cashout
	cashOutRepository := cashoutRepoPkg.NewCashOutRepository(config.DB)
	cashOutUC := cashoutUsecasePkg.NewCashOutUsecase(cashOutRepository)
	cashOutH := cashoutHandlerPkg.NewCashOutHandler(cashOutUC)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth.RegisterAuthRoutes(r, authH)

	r.Static("/uploads", "./uploads")

	protected := r.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		auth.RegisterMeRoute(protected)
		master.RegisterMasterSizeRoutes(protected, masterSizeH)
		master.RegisterMasterMaterialRoutes(protected, masterMaterialH)
		master.RegisterMasterSaintsRoutes(protected, masterSaintsH)
		master.RegisterMasterStyleRoutes(protected, masterStyleH)
		master.RegisterMaterialComponentRoutes(
			protected,
			materialComponentH,
		)
		master.RegisterMaterialComponentVariantRoutes(
			protected,
			materialComponentVariantH,
		)
		transaction.RegisterClientRoutes(protected, clientH)
		commission.RegisterCommissionRoutes(protected, commissionH)
		commission.RegisterCommissionItemMaterialRoutes(
			protected,
			commissionItemMaterialHandler,
		)
		cashout.RegisterCashOutRoutes(protected, cashOutH)
		protected.POST("/upload", uploadHandler.UploadFile)
	}

	r.Run(":8080")
}
