package models

import "gorm.io/gorm"

type BuyItem struct {
	gorm.Model
	Name string `json:"name" gorm:"text;not null;default:null`
}
