package model

import (
	"todo-rest-api/database"

	"gorm.io/gorm"
)

type Todo struct {
    gorm.Model
    Content string `gorm:"type:text" json:"content"`
    UserID  uint
}

func (entry *Todo) Save() (*Todo, error) {
    err := database.Database.Create(&entry).Error
    if err != nil {
        return &Todo{}, err
    }
    return entry, nil
}