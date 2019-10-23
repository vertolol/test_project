package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
)


type PostgresWorker struct {
	DB 			*gorm.DB
}

type ConnectionConfig struct {
	USER 		string
	PASSWORD 	string
	DB_NAME 	string
	HOST 		string
	PORT 		string
}


func CreatePostgresConnection(config *ConnectionConfig) (db *gorm.DB, err error) {
	url := fmt.Sprintf("user=%s password=%s DB.name=%s host=%s port=%s sslmode=disable",
		config.USER,
		config.PASSWORD,
		config.DB_NAME,
		config.HOST,
		config.PORT,
	)

	db, err = gorm.Open("postgres", url)
	if err != nil {
		return db, err
	}

	return db, err
}
