package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Task comment
type Task struct {
	gorm.Model
	Name string `json:"name"`
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func allTasks(rw http.ResponseWriter, req *http.Request) {

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	tasks := []Task{}
	db.Find(&tasks)
	respondJSON(rw, http.StatusOK, tasks)
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

	r := mux.NewRouter()

	r.HandleFunc("/tasks", allTasks)

	log.Fatal(http.ListenAndServe(":"+port, r))
}
