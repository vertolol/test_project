package main

import (
	"api/api_routes"
	"api/app"
	"api/elastic"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)


func main() {
	var Config app.TomlConfig
	_, err := toml.DecodeFile("config.toml", &Config)
	app.IfError(err)

	client, err := elastic.CreateElasticConnection(&Config.Elastic)
	app.IfError(err)

	elasticClient := &elastic.ElasticWorker{
		Client: client,
		Index: Config.Elastic.INDEX_NAME,
	}

	App_connections := &app.App{
		Elastic: *elasticClient,
		Config: Config,
	}

	r := gin.Default()
	api_routes.InitializeRoutes(r, App_connections)

	r.Run(":3001")
}
