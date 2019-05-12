package main

import (
	_ "todoapp-api/data/utils"
	"log"
	"net/http"
	"os"

	"todoapp-api/data/repository"

	"github.com/gin-gonic/gin"
)

func allTasks(c *gin.Context) {
	//respondJSON(rw, http.StatusOK, tasks)
	c.JSON(http.StatusOK, repository.GetAllTasks())
}

//http://www.golangprograms.com/golang-restful-api-using-grom-and-gorilla-mux.html
func main() {
	// Heroku supplies your port via environment variable
	port := os.Getenv("PORT")
	r := gin.Default()

	r.GET("/tasks", allTasks)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
