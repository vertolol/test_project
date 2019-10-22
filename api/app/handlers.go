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

	db, err := postgres.CreatePostgresConnection(&App.Config.Postgres)
	IfError(err)

	postgresWorker := &postgres.PostgresWorker{DB: db}

	var result []postgres.Product
	for _, id := range ids {
		product := postgresWorker.GetInstanceById(id)
		result = append(result, product)
	}

	postgresWorker.DB.Close()

	c.JSON(
		http.StatusOK,
		gin.H{
			"products": result,
		},
	)
}

