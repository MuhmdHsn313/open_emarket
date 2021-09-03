package routes

import (
	"github.com/gin-gonic/gin"
	v1 "open_emarker/controllers/v1"
	"open_emarker/services/authentication/jwt"
)

func RegisterUserRoutes(versionGroup *gin.RouterGroup) {
	c_user := v1.UserController{}

	group := versionGroup.Group("/user")
	{
		group.POST("/create", c_user.CreateAccount)
		group.POST("/login", c_user.LoginToAccount)
		group.POST("/me", jwt.JwtAuthorization(), c_user.GetMyAccount)
	}
}
