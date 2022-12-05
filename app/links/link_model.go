package links

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	ID           string         `json:"id" gorm:"primaryKey"`
	LinkID       string         `json:"link_id" gorm:"uniqueIndex"`
	OriginalLink string         `json:"original_link" gorm:"not null"`
	ShortLink    string         `json:"short_link" gorm:"not null"`
	ApiKey       string         `json:"api" gorm:"not null"`
	Type         string         `json:"type" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeleteAt     gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
