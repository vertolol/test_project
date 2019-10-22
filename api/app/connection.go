package app


import (
	"api/elastic"
	"api/postgres"
)


type TomlConfig struct {
	Elastic   	elastic.ConnectionConfig
	Postgres  	postgres.ConnectionConfig
}


type App struct {
	Elastic 	elastic.ElasticWorker
	Postgres 	postgres.PostgresWorker
	Config 		TomlConfig
}
