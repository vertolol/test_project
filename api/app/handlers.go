package app


import (
	"api/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)


func (App *App) SearchProduct(c *gin.Context) {
	name := c.Query("name")

	ids, err := App.Elastic.GetIdsByName(name)
	IfError(err)

	var result []postgres.Product
	for _, id := range ids {
		product := App.Postgres.GetInstanceById(id)
		result = append(result, product)
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"products": result,
		},
	)
}
