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
		config.GetEnvVarAsString("DB_USER", "ryan"),
		config.GetEnvVarAsString("DB_PASSWORD", "1234"),
		config.GetEnvVarAsString("DB_NAME", "go_gin_example"),
		config.GetEnvVarAsString("DB_PORT", "5432"),
		config.GetEnvVarAsString("DB_SSLMODE", "disable"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	// enum타입 생성
	db.Exec("CREATE TYPE IF NOT EXISTS wish_status AS ENUM ('approved', 'pending', 'rejected');")
	// AutoMigrate는 테이블이 없으면 생성하고, 필드가 없으면 추가함
	db.AutoMigrate(&models.Wish{})
	// db.AutoMigrate(&models.Comment{})
	return db
}
