package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Url string
}
