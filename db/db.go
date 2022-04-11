package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() (err error) {
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		err := os.Mkdir("./data", os.ModePerm)
		if err != nil {
			return err
		}
	}
	DB, err = gorm.Open(sqlite.Open("./data/sqlite3.db"), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		return err
	}
	return err
}
