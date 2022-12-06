package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    string         `json:"user_id"`
	ApiKey    string         `json:"apikey" gorm:"unique;not null;index"`
	Email     string         `json:"email"`
	Type      string         `json:"type" gorm:"not null"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeleteAt  gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
