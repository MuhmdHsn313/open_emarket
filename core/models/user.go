package models

import (
	"database/sql"
	"open_emarker/lib/encryption"
	"open_emarker/settings"
	"time"
)

type User struct {
	Id          uint         `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"not null;size:255"`
	Email       *string      `json:"email" gorm:"not null; type:varchar(100);unique"`
	Username    *string      `json:"username" gorm:"not null;type:varchar(100);unique"`
	PhoneNumber string       `json:"phone_number" gorm:"not null"`
	Password    string       `json:"password" gorm:"not null" `
	Photo       string       `json:"photo" gorm:""`
	LastLogin   sql.NullTime `json:"last_login"`
	CreatedDate time.Time    `json:"created_date" gorm:"not null;autoCreateTime"`
	Birthday    *time.Time   `json:"birthday" gorm:"not null"`
}

func (user User) SetPassword(password string) {

	encryptPassword, err := encryption.Encrypt([]byte(password))
	if err != nil {
		panic("Can't set password!")
	}

	settings.DB.Model(&user).Update("password", string(encryptPassword))
}

func (user User) CheckPassword(password string) bool {
	realPassword, err := encryption.Decrypt([]byte(user.Password))
	if err != nil {
		return false
	}

	return string(realPassword) == password
}
