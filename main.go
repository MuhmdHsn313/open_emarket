package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	routes "open_emarker/routers/v1"
	"open_emarker/services/config"
	"open_emarker/services/sqlite"
	"open_emarker/settings"
)

func main() {

	fmt.Println("Welcome!\nCheck Configuration...")
	configure, err := config.SetupServerConfigure()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the welcome message
	fmt.Println("Please wait to configure the database!")

	// Setup database database
	settings.DB = sqlite.SetupDatabase(configure.DSN)

	// Print start app message
	fmt.Println("Configurations done! app starting...")

	// Setup publishing mode
	gin.SetMode(gin.DebugMode)

	// Setup webserver
	server := gin.Default()

	// Set settings
	server.MaxMultipartMemory = settings.MaxMultipartMemory

	// Register routes callable
	routes.RegisterVersion1Routes(server)

	// Run the server
	err = server.Run(configure.GetFullAddress())
	if err != nil {
		fmt.Println(err)
		return
	}

}
