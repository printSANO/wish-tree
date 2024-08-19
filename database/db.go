package database

import (
	"fmt"
	"log"

	"github.com/printSANO/wish-tree/config"
	"github.com/printSANO/wish-tree/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	config.LoadEnvFile(".env")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.GetEnvVarAsString("DB_HOST", "localhost"),
		config.GetEnvVarAsString("DB_USER", "techeer"),
		config.GetEnvVarAsString("DB_PASSWORD", "secret"),
		config.GetEnvVarAsString("DB_NAME", "hackathon"),
		config.GetEnvVarAsString("DB_PORT", "5432"),
		config.GetEnvVarAsString("DB_SSLMODE", "disable"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	// enum타입 wish_status를 생성
	db.Exec("CREATE TYPE wish_status AS ENUM ('approved', 'pending', 'rejected');")
	log.Println("Create wish_status enum type")

	// AutoMigrate는 테이블이 없으면 생성하고, 필드가 없으면 추가함
	db.AutoMigrate(&models.Wish{})
	log.Println("AutoMigrate wish table")

	// db.AutoMigrate(&models.Comment{})
	// log.Println("AutoMigrate comment table")
	return db
}
