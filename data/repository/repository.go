package repository

import (
	"log"
	"os"
	"todoapp-api/dto"

	"github.com/jinzhu/gorm"
)

// GetAllTasks devuelve todas las tareas
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

func DeleteTask(id int) bool {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	task, err := getTaskOrNil(db, id)
	if task.ID == 0 {
		log.Println("task doesnt exist")
		return false
	}

	if err != nil {
		log.Fatal(err)
	}


	err = db.Delete(task).Error
	if err != nil {
		log.Fatal(err)
	}

	return true
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

