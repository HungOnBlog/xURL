package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId          string         `json:"user_id"`
	ApiKey          string         `json:"apikey" gorm:"unique;not null;index"`
	Email           string         `json:"email"`
	Type            string         `json:"type" gorm:"not null"`
	Name            string         `json:"name"`
	Password        string         `json:"password"`
	LimitPassword   int            `json:"limit_password" gorm:"default:0"`
	LimitTracking   int            `json:"limit_tracking" gorm:"default:0"`
	LimitCustomSlug int            `json:"limit_custom_slug" gorm:"default:0"`
	CreatedAt       time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeleteAt        gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
