package routes

import "github.com/gin-gonic/gin"

func RegisterVersion1Routes(server *gin.Engine) {
	v1 := server.Group("v1")
	RegisterUserRoutes(v1)
}
