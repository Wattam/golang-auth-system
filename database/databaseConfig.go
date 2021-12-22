package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const host = "localhost"
const port = "5432"
const user = "postgres"
const password = "postgres123"
const dbName = "ShoeStore"

var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

var Db *gorm.DB

func ConnectDatabase() {

	dbGorm, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database connected!")
	}

	Db = dbGorm

	RunMigrations()
}

func RunMigrations() {
	Db.AutoMigrate()
}

func DisconnectDatabase() {
	config, _ := Db.DB()

	config.Close()
}
