package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	_ "todoapp-api/data/utils"
	"todoapp-api/dto"

	"todoapp-api/data/repository"

	"github.com/gin-gonic/gin"
)

func allTasks(c *gin.Context) {
	//respondJSON(rw, http.StatusOK, tasks)
	c.JSON(http.StatusOK, repository.GetAllTasks())
}

func newTask(c *gin.Context) {
	var json dto.Task
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, repository.CreateTask(json))
}

func modifyTask(c *gin.Context) {
	var json dto.UpdateFields
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	idnumber, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, repository.UpdateTask(idnumber, json.Name))
}

func deleteTask(c *gin.Context) {

	id := c.Param("id")

	idnumber, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, repository.DeleteTask(idnumber))
}

//http://www.golangprograms.com/golang-restful-api-using-grom-and-gorilla-mux.html
func main() {
	// Heroku supplies your port via environment variable
	port := os.Getenv("PORT")
	r := gin.Default()

	r.GET("/tasks", allTasks)
	r.POST("tasks", newTask)
	r.PUT("tasks/:id", modifyTask)
	r.DELETE("tasks/:id",deleteTask)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
