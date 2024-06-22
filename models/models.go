package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	ID            uint   `gorm:"primary_key"`
	Title         string `gorm:"column:name"`
	GameID        int32
	GameURL       string
	IsCooperative bool
}

type User struct {
	gorm.Model
	ID     uint   `gorm:"primary_key"`
	Name   string `gorm:"column:name"`
	UserID string
}
