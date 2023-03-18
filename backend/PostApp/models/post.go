package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	// ID          int       `json:"id"`
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURI    string `json:"imageuri"`
}
