package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events")
	server.GET("/events")
	server.POST("/events")
	server.PUT("/events")
}
