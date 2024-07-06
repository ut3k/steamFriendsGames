package main

import (
	// "path/filepath"

	"os"
	// "steamFriendsGames/checkcoop"
	// "steamFriendsGames/controllers"
	"steamFriendsGames/generators"
	"steamFriendsGames/initialisers"
)

func init() {
	initialisers.ConnectToDataBase()
	initialisers.SyncDB()
}

func main() {
	// controllers.ScrapeLocalData()
	// controllers.CheckIfGameHasManyUsers()
	// checkcoop.CheckIfGameIsCoop()
	generators.GenerateHTMLlist()
	os.Exit(1)
}
