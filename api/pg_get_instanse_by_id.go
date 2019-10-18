package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func getInstanceById(id int64) Product{
	const (
		user     = "gorm"
		password = "gorm"
		dbname   = "gorm"
		host     = "postgres_host"
		port     = "5432"
	)

	url := fmt.Sprintf("user=%s password=%s DB.name=%s host=%s port=%s sslmode=disable",
		user,
		password,
		dbname,
		host,
		port,
	)

	db, err := gorm.Open("postgres", url)
	defer db.Close()
	ifPanic(err)

	var product Product
	db.Find(&product, id)

	return product
}
