package infrastructure

import (
	"log"
	"os"

	"github.com/Ateto1204/swep-user-serv/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	godotenv.Load()
	dsn := os.Getenv("POSTGRESQL_CONNECTION")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
		return nil, err
	} else {
		log.Println("Migrated database successfully")
	}
	return db, nil
}
