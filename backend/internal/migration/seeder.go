package migration

import (
	"log"

	admin "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/admin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	var count int64
	db.Model(&admin.Admin{}).Count(&count)

	if count > 0 {
		log.Println("Admin sudah ada, skip seeding")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Gagal hash password seeder: ", err)
	}

	adminData := admin.Admin{
		Username: "admin",
		Password: string(hashedPassword),
	}

	if err := db.Create(&adminData).Error; err != nil {
		log.Fatal("Gagal seed admin: ", err)
	}

	log.Println("Seeder berhasil: admin/admin123")
}