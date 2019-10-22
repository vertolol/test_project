package main

import (
	"github.com/gin-gonic/gin"
)


type Env struct {
	elastic ElasticWorker
}


func main() {
	r := gin.Default()

	elasticClient := &ElasticWorker{client: createElasticConnection()}
	env := &Env{elastic: *elasticClient}

	r.GET("/search", env.searchProduct)

	r.Run(":3001")
}
