package models

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Id           uint           `json:"id" gorm:"primaryKey"`
	OriginalLink string         `json:"originalLink"`
	LinkId       string         `json:"linkId"`
	ApiKey       string         `json:"apiKey"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
}
