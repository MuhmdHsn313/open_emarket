package main

import (
	"github.com/gin-gonic/gin"
	"open_emarker/services/config"
	"open_emarker/services/sqlite"
	"open_emarker/settings"
)

func main() {

	println("Welcome!\nCheck Configuration...")
	configure, err := config.SetupServerConfigure()
	if err != nil {
		println(err.Error())
		return
	}

	// Print the welcome message
	println("Welcome!\nPlease wait to configure the database!")

	// Setup database database
	settings.DB = sqlite.SetupDatabase(configure.DSN)

	// Print start app message
	println("Configurations done! You can start and use app.")

	// Setup publishing mode
	gin.SetMode(gin.DebugMode)

	// Setup webserver
	server := gin.Default()

	// Run the server
	err = server.Run(configure.GetFullAddress())
	if err != nil {
		return
	}

}
