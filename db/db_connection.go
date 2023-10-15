package db

import (
	"fmt"
	"go_algo/db/connect"
	"go_algo/db/models"
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB = nil

func Connect() {

	var err error

	db, err = connect.GetDB()

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("database connected successfully")
	}

	// Migrate all the models.
	err = db.AutoMigrate(&models.Student{})

	if err != nil {
		log.Fatal("unable to create schema", err)
	} else {
		println("schemas created successfully")
	}
}
