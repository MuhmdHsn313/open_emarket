package main

import (
	"github.com/gin-gonic/gin"
	"open_emarker/services/config"
)

func main() {

	println("Welcome!\nCheck Configuration...")
	configure, err := config.SetupServerConfigure()
	if err != nil {
		println(err.Error())
		return
	}

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
