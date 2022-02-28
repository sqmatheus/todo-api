package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TodoModel struct {
	ID        string `gorm:"primaryKey;"`
	Title     string `gorm:"type:varchar(255);unique;not null"`
	Done      bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *TodoModel) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.UpdatedAt = time.Now().Local()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *TodoModel) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
