package utils

import (
	"log"
	"os"
	"todoapp-api/data/model"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)

	//migrations
	db.AutoMigrate(&model.Task{})
}
