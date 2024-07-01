package main

import (
	// "os"
	// "path/filepath"

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
	// controllers.CheckIfGameIsCoop()
}
