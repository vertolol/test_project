package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func (client *ElasticWorker) searchProduct(c *gin.Context) {
	name := c.Query("name")

	ids := getIdsByName(name, client.client)
	db := createPostgresConnection()

	var result []Product
	for _, id := range ids {
		product := getInstanceById(id, db)
		result = append(result, product)
	}

	db.Close()

	c.JSON(
		http.StatusOK,
		gin.H{
			"products": result,
		},
	)
}

