package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jaseelaali/passwordgenerator_unique/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() {
	// load database

	if err := godotenv.Load(); err != nil {
		fmt.Println("error in loading env file")
	}
	// connect database

	dsn := os.Getenv("DB")
	fmt.Println("dsn:", dsn)
	var err error
	DB, err = gorm.Open(postgres.Open("user=postgres host=localhost port=5432 password=password dbname=example sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("error in connecting database")
	}
	log.Println("Successfully connected to database")
	//sync database
	err = DB.AutoMigrate(
		&model.Users{},
		&model.Signup{},
	)

}
