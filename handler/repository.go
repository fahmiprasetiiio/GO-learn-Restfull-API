package handler

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Ini untuk DB
)

func getConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", os.Getenv("DBCONN"))
	if err != nil {
		return
	}

	return
}
