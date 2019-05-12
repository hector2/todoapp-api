package repository

import (
	"log"
	"os"
	"todoapp-api/dto"

	"github.com/jinzhu/gorm"
)

func GetAllTasks() []dto.Task {

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	tasks := []dto.Task{}
	db.Find(&tasks)
	return tasks
}

func CreateTask(task dto.Task) dto.Task {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	db.Create(&task)

	return task
}

