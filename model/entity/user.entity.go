package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `db:"id" json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
