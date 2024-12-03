package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luizenmattos/go_api/handler"
)

const Version = "v1"

func initializeRoutes(router *gin.Engine) {
	routes := router.Group("/api/" + Version)
	{
		routes.GET("/opening", handler.ShowOpeningHandler)
		routes.GET("/openings", handler.ListOpeningsHandler)
		routes.PUT("/opening", handler.UpdateOpeningHandler)
		routes.POST("/opening", handler.CreateOpeningHandler)
		routes.DELETE("/opening", handler.DeleteOpeningHandler)
	}
}
