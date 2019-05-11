package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Task comment
type Task struct {
	gorm.Model
	Name string `json:"name"`
}

func allTasks(c *gin.Context) {


	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	tasks := []Task{}
	db.Find(&tasks)
	//respondJSON(rw, http.StatusOK, tasks)
	c.JSON(http.StatusOK,tasks)
}

func InitDB() {
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
	db.AutoMigrate(&Task{})
}

//http://www.golangprograms.com/golang-restful-api-using-grom-and-gorilla-mux.html
func main() {
	// Heroku supplies your port via environment variable
	port := os.Getenv("PORT")

	InitDB()
	r := gin.Default()


	r.GET("/tasks", allTasks)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
