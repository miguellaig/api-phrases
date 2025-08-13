package models

import "gorm.io/gorm"

type Phrases struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Original    string `json:"orignial"`
	Translation string `json:"translation"`
	Language    string `json:"language"`
}
