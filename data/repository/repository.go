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

func UpdateTask(id int, name string) dto.Task {

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	task, err := getTaskOrNil(db, id)

	if err != nil {
		log.Fatal(err)
	}

	task.Name = name
	db.Save(task)
	return task
}

func getTaskOrNil(db *gorm.DB, id int) (dto.Task,error) {
	task := dto.Task{}
	return task, db.First(&task, id).Error
}

