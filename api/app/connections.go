package app

import (
	"api/elastic"
	"api/postgres"
)


type App struct {
	Elastic 	elastic.ElasticWorker
	Postgres 	postgres.PostgresWorker
}
