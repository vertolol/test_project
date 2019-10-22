package api_routes

import (
	"api/app"
	"github.com/gin-gonic/gin"
)


func InitializeRoutes(r *gin.Engine, AppConnections *app.App) {

	r.GET("/search", AppConnections.SearchProduct)

}