package main

import (
	"api/api_routes"
	"api/app"
	"api/elastic"
	"api/postgres"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)


type TomlConfig struct {
	Elastic   	elastic.ConnectionConfig
	Postgres  	postgres.ConnectionConfig
}


func main() {
	var Config TomlConfig
	_, err := toml.DecodeFile("config.toml", &Config)
	app.IfError(err)

	elastic_connection, err := elastic.CreateElasticConnection(&Config.Elastic)
	app.IfError(err)
	elasticClient := &elastic.ElasticWorker{
		Client: elastic_connection,
		Index: Config.Elastic.INDEX_NAME,
	}

	postgres_connection, err := postgres.CreatePostgresConnection(&Config.Postgres)
	app.IfError(err)
	postgresDB := &postgres.PostgresWorker{
		DB:postgres_connection,
	}

	App_connections := &app.App{
		Elastic: *elasticClient,
		Postgres: *postgresDB,
	}

	r := gin.Default()
	api_routes.InitializeRoutes(r, App_connections)

	r.Run(":3001")
}
