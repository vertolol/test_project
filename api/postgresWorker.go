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


func (postgres PostgresWorker) getInstanceById(id int64) Product{
	var product Product
	postgres.db.Find(&product, id)

	return product
}
