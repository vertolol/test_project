package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func (env *Env) searchProduct(c *gin.Context) {
	name := c.Query("name")

	ids := env.elastic.getIdsByName(name)
	postgres := &PostgresWorker{db: createPostgresConnection()}

	var result []Product
	for _, id := range ids {
		product := postgres.getInstanceById(id)
		result = append(result, product)
	}

	postgres.db.Close()

	c.JSON(
		http.StatusOK,
		gin.H{
			"products": result,
		},
	)
}

