package routes

import (
	"net/http"
	model "ticketsystem/models"

	"github.com/gin-gonic/gin"
)

func getAllTickets(context *gin.Context) {
	tickets := model.GetAllTickets()
	context.JSON(200, gin.H{"message": "This is a GET request", "tickets": tickets})
}

func createTicket(context *gin.Context) {
	var ticket model.Ticket
	err := context.ShouldBindJSON(&ticket)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "There was some error parsing the ticket", "err": err.Error()})
	}
	ticket.ID = 1
	ticket.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "New ticket created", "ticket": ticket})

}
