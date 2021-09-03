package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"open_emarker/core/models"
	"open_emarker/settings"
)

func JwtAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := InstanceFromJwtService().ValidateToken(tokenString)
		if token.Valid {
			var user models.User
			if id, ok := token.Claims.(jwt.MapClaims)["id"]; ok {
				err = settings.DB.Where("id = ?", id.(string)).First(&user).Error
				if err != nil {
					fmt.Println(err)
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You do not have authentication!"})
					return
				}
				ctx.Set("User", &user)
			} else {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You do not have authentication!"})
			}

		} else {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You do not have authentication!"})
		}
	}
}
