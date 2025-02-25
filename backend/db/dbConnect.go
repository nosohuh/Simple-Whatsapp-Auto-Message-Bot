package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	usermodels "github.com/while/payproje/models/UserModels"
	servis "github.com/while/payproje/models/ServisModels"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	DbHost := os.Getenv("POSTGRESQL_HOST")
	DbName := os.Getenv("POSTGRESQL_DBNAME")
	DbUsername := os.Getenv("POSTGRESQL_USER")
	DbPassword := os.Getenv("POSTGRESQL_PASSWORD")
	DbPort := os.Getenv("POSTGRESQL_PORT")

	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DbHost, DbPort, DbUsername, DbName, DbPassword)
	dbConnection, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic("database connection failed")
	}
	DB = dbConnection

	AutoMigrate(dbConnection)

	fmt.Println("database connected successfully")

}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		//DEALER MODELS
		&servis.Servis{},
		//DEPOSIT MODELS
		&usermodels.Users{},
	)
}
