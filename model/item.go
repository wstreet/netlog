package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Title     string `json:"title"`
	OriginUrl string `json:"originUrl"`
	Name      string `json:"name"`
	Href      string `json:"href"`
}
