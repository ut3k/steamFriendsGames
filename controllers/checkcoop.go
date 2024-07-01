package controllers

// import (
// 	"fmt"
// "log"
// "net/http"
// "os"
// "path/filepath"
// "strings"

// "steamFriendsGames/models"

// "github.com/gocolly/colly"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// func CheckIfGameIsCoop() {
// 	var err error
// 	var DB *gorm.DB
// 	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("Function CheckIfGameIsCoop can not connect to data.db")
// 	}

// 	var games []models.Game
// 	err = DB.Preload("usergames").Where("user_id > ?", 1).Find(&games).Error
// 	if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println(err)
// 	}
// }
