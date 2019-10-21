package main

import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()

	elasticClient := &ElasticWorker{client: createElasticConnection()}
	r.GET("/search", elasticClient.searchProduct)
	//initializeRoutes(r)

	r.Run(":3001")
}
