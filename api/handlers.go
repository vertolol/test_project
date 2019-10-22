package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func (client *ElasticWorker) searchProduct(c *gin.Context) {
	name := c.Query("name")

	ids := client.getIdsByName(name)
	db := &PostgresWorker{db: createPostgresConnection()}

	var result []Product
	for _, id := range ids {
		product := db.getInstanceById(id)
		result = append(result, product)
	}

	db.db.Close()

	c.JSON(
		http.StatusOK,
		gin.H{
			"products": result,
		},
	)
}

