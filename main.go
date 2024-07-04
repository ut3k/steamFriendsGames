package main

import (
	// "os"
	// "path/filepath"

	"steamFriendsGames/checkcoop"
	"steamFriendsGames/controllers"
	"steamFriendsGames/initialisers"
)

func init() {
	initialisers.ConnectToDataBase()
	initialisers.SyncDB()
}

func main() {
	controllers.ScrapeLocalData()
	controllers.CheckIfGameHasManyUsers()
	checkcoop.CheckIfGameIsCoop()
}
