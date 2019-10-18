package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func searchProduct(c *gin.Context) {
	name := c.Query("name")
	ids := getIdsByName(name)

	var res []Product
	for _, id := range ids {
		product := getInstanceById(id)
		res = append(res, product)
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"products": res,
		},
	)
}

