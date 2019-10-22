package postgres

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func (postgres PostgresWorker) GetInstanceById(id int64) Product{
	var product Product
	postgres.DB.Find(&product, id)

	return product
}
