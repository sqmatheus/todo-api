package database

import (
	"github.com/sqmatheus/todo-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Initialize(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		defer log.Fatal("ERROR: could not connect to database")
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&models.TodoModel{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("INFO: successfully connected to database")

	return db
}
