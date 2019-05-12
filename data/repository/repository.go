package repository

import (
	"log"
	"os"
	"todoapp-api/data/model"
	"github.com/jinzhu/gorm"
)

func GetAllTasks() []model.Task {

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	tasks := []model.Task{}
	db.Find(&tasks)
	return tasks
}
