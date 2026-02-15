// Copyright by AcmaTvirus
package database

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	dbPath := "data/foxdocker.db"
	os.MkdirAll(filepath.Dir(dbPath), 0755)

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto Migration
	log.Println("Database migration started...")
	return DB.AutoMigrate(&Project{}, &User{})
}

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
}
