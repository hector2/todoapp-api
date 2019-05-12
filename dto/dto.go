package dto

import (
	"github.com/jinzhu/gorm"
)

//Task comment
type Task struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}
