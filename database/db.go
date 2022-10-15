package database

import (
	"fmt"
	"test/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbName   = "orders_by"
)

var (
	db  *gorm.DB
	err error
)

func Connect(migrateDB bool) (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, dbPort, dbName)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if migrateDB {
		db.Debug().AutoMigrate(models.Order{}, models.Item{})
	}
	return db, nil
}

func CloseDB() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
