package models

import (
	"database/sql"
	"open_emarker/lib/encryption"
	"open_emarker/settings"
	"strings"
	"time"
)

type User struct {
	Id          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;type:varchar(150)"`
	Email       *string        `json:"email" gorm:"not null; type:varchar(100);unique"`
	Username    *string        `json:"username" gorm:"not null;type:varchar(100);unique"`
	PhoneNumber string         `json:"phone_number" gorm:"not null;type:varchar(50)"`
	Password    string         `json:"password" gorm:"not null"`
	Photo       sql.NullString `json:"photo" gorm:"type:varchar(255)"`
	LastLogin   sql.NullTime   `json:"last_login"`
	CreatedDate time.Time      `json:"created_date" gorm:"not null;autoCreateTime"`
	Birthday    *time.Time     `json:"birthday" gorm:"not null"`
}

func (user *User) SetPassword(password string) {

	encryptPassword := user.GetEncryptedPassword(password)

	user.Password = encryptPassword
	settings.DB.Model(&user).Update("password", encryptPassword)

}

func (user User) GetEncryptedPassword(password string) string {

	encryptPassword, err := encryption.Encrypt([]byte(password))
	if err != nil {
		panic("Can't set password!")
	}
	return string(encryptPassword)

}

func (user *User) CheckPassword(password string) bool {

	realPassword, err := encryption.Decrypt([]byte(user.Password))
	if err != nil {
		return false
	}

	return string(realPassword) == password

}

func (user User) GetDataShown(args ...interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"id":           user.Id,
		"name":         user.Name,
		"email":        user.Email,
		"username":     user.Username,
		"phone_number": user.PhoneNumber,
		"birthday":     user.Birthday,
		"created_date": user.CreatedDate,
	}

	if len(args)%2 == 0 {
		for i := 0; i < len(args); i += 2 {
			data[args[i].(string)] = args[i+1]
		}
	}

	if user.LastLogin.Valid {
		data["last_login"] = user.LastLogin.Time
	} else {
		data["last_login"] = nil
	}

	if user.Photo.Valid {
		data["photo"] = user.Photo.String
	} else {
		data["photo"] = nil
	}

	return data
}

// CreateAccountInput using when create a new user account.
type CreateAccountInput struct {
	Name        string     `json:"name" xml:"name" binding:"required"`
	Email       *string    `json:"email" xml:"email" binding:"required"`
	Username    *string    `json:"username" xml:"username" binding:"required"`
	PhoneNumber string     `json:"phone_number" xml:"phone_number" binding:"required"`
	Password    string     `json:"password" xml:"password" binding:"required"`
	Birthday    *time.Time `json:"birthday" xml:"birthday" binding:"required"`
	Photo       string     `json:"photo" xml:"photo"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (i LoginInput) IsEmail() bool {
	return strings.Contains(i.Username, "@")
}
