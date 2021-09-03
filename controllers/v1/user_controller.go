package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"open_emarker/core/models"
	"open_emarker/services/authentication/jwt"
	"open_emarker/settings"
	"time"
)

type UserController struct {
	server *gin.Engine
}

func (c UserController) CreateAccount(ctx *gin.Context) {

	var createInput models.CreateAccountInput
	err := ctx.ShouldBindJSON(&createInput)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:        createInput.Name,
		Email:       createInput.Email,
		Username:    createInput.Username,
		PhoneNumber: createInput.PhoneNumber,
		LastLogin:   sql.NullTime{},
		CreatedDate: time.Time{},
		Birthday:    createInput.Birthday,
	}
	user.Password = user.GetEncryptedPassword(createInput.Password)

	settings.DB.Create(&user)

	token := jwt.InstanceFromJwtService().GenerateToken(user.Id)
	ctx.JSON(http.StatusCreated, user.GetDataShown("token", token))

}

func (c UserController) LoginToAccount(ctx *gin.Context) {

	var loginInput models.LoginInput
	if err := ctx.ShouldBindJSON(&loginInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cond string
	var user models.User

	if loginInput.IsEmail() {
		cond = "email = ?"
	} else {
		cond = "username = ?"
	}

	err := settings.DB.Where(cond, loginInput.Username).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Please verify that your login information is correct!"})
		return
	}

	if user.CheckPassword(loginInput.Password) {
		token := jwt.InstanceFromJwtService().GenerateToken(user.Id)
		user.UpdateLastLoginTime()
		ctx.JSON(http.StatusOK, user.GetDataShown("token", token))
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Please verify that your login information is correct!"})
	}

}

func (c UserController) GetMyAccount(ctx *gin.Context) {
	user := ctx.MustGet("User").(*models.User)
	ctx.JSON(http.StatusOK, user.GetDataShown())
}

func (c UserController) RefreshToken(ctx *gin.Context) {
	user := ctx.MustGet("User").(*models.User)
	token := jwt.InstanceFromJwtService().GenerateToken(user.Id)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c UserController) UploadPhoto(ctx *gin.Context) {

	photo, err := ctx.FormFile("photo")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := ctx.MustGet("User").(*models.User)
	dst := settings.GetImagePath(photo)
	err = ctx.SaveUploadedFile(photo, dst)
	user.UpdateImage(dst)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user.GetDataShown())
}

func (c UserController) UpdateAccount(ctx *gin.Context) {

	var updateInput models.UpdateAccountInput
	err := ctx.ShouldBindJSON(&updateInput)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := ctx.MustGet("User").(*models.User)
	user.UpdateWith(updateInput)
	ctx.JSON(http.StatusOK, user.GetDataShown())

}
