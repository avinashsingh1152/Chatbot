package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	IPAddress string `gorm:"column:ip_address"`
	ImageURL  string `gorm:"column:image_url"`
	FileName  string `gorm:"column:file_name"`
}

func (i Image) TableName() string {
	return "images"
}
