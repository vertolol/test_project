package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


type PostgresWorker struct {
	db *gorm.DB
}


func createPostgresConnection() *gorm.DB {
	url := fmt.Sprintf("user=%s password=%s DB.name=%s host=%s port=%s sslmode=disable",
		POSTGRES_USER,
		POSGRES_PASSWORD,
		POSTGRES_DB_NAME,
		POSTGRES_HOST,
		POSTGRES_PORT,
	)

	db, err := gorm.Open("postgres", url)
	ifPanic(err)
	return db
}


func getInstanceById(id int64, db *gorm.DB) Product{
	var product Product
	db.Find(&product, id)

	return product
}
