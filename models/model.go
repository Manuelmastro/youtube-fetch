package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Title        string    `gorm:"column:title" json:"title"`
	Description  string    `gorm:"column:description" json:"description"`
	PublishTime  time.Time `gorm:"column:publish_datetime" json:"publish_datetime"`
	ThumbnailURL string    `gorm:"column:thumbnail_url" json:"thumbnail_url"`
	VideoID      string    `gorm:"column:video_id;uniqueIndex" json:"video_id"` // Unique index for VideoID
	SearchVector string    `gorm:"column:search_vector;type:tsvector;index:search_vector_idx,gin;->" json:"-"`
}
