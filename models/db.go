package models

import (
	config "github.com/rorikurniadi/simple-auth/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// InitDB creates and migrates the database
func InitDB() (*gorm.DB, error) {
	var err error

	var dbConfig = config.ReadConfig()

	dbUser := dbConfig.DB_USER
	dbPassword := dbConfig.DB_PASSWORD
	dbName := dbConfig.DB_NAME
	connectionString := dbUser + ":" + dbPassword + "@/" + dbName + "?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	db.AutoMigrate(&User{})

	return db, nil
}
