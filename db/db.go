package db

import (
	"fmt"
	"log"
	"os"
	"todo-api/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := gorm.Open(sqlite.Open("test"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.User{}, &model.Task{})

	fmt.Println("DB connection successfully opened")
	return db
}
