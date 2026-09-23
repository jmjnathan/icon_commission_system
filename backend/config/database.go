package config

import (
	"fmt"
	"log"
	"os"

	admin "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	cashoutEntity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/cashflow/cashout"
	client "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/client"
	commissionEntity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/commission"
	master "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/master"
	"github.com/nathanchristiawan02/icon-commission-system-backend/internal/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database: ", err)
	}

	log.Println("Database berhasil terkoneksi!")
	err = db.AutoMigrate(
		&admin.Admin{},
		&master.MasterSize{},
		&master.MasterMaterial{},
		&master.MaterialComponent{},
		&master.MaterialComponentVariant{},
		&master.MasterSaints{},
		&master.MasterStyle{},
		&client.Client{},
		&commissionEntity.Commission{},
		&commissionEntity.CommissionItem{},
		&commissionEntity.CommissionItemMaterial{},
		&commissionEntity.CommissionPhoto{},
		&commissionEntity.CommissionPayment{},
		&cashoutEntity.CashOut{},
	)
	if err != nil {
		log.Fatal("Gagal migrate: ", err)
	}
	log.Println("Migration berhasil!")

	migration.SeedAdmin(db)

	DB = db
}
