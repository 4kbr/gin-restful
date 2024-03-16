package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	// https://github.com/go-gorm/postgres
	dsn := os.Getenv("DB_URI")
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")

	}
	fmt.Println("successfully connect db")
}
