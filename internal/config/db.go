package config

import (
	"log"
	"os"

	"github.com/msyahruls/kreditplus-go-test/internal/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Run Auto Migration
	err = db.AutoMigrate(&domain.User{}, &domain.Limit{}, &domain.Transaction{}, &domain.PaymentSchedule{})
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	return db
}
