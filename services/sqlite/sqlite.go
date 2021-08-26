package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"open_emarker/core/models"
)

func ConfigureNewDB(dbName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		panic("Error with configure database!")
	}

	return db
}

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Error with migrate the new database.")
	}
}

func SetupDatabase(dbName string) *gorm.DB {

	db := ConfigureNewDB(dbName)

	Migration(db)

	return db

}
