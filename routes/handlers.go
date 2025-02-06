package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/tickets", getAllTickets)
	server.POST("/tickets", createTicket)
	server.GET("/tickets/:id", getSingleTicket)
	server.PUT("/tickets/:id", updateTicket)

}
