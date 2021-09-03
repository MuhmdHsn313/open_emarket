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
		group.GET("/refresh", jwt.JwtAuthorization(), c_user.RefreshToken)
		me := group.Group("/me").Use(jwt.JwtAuthorization())
		me.GET("", c_user.GetMyAccount)
		me.PATCH("", c_user.UpdateAccount)
		me.PUT("/photo", c_user.UploadPhoto)
	}
}
