package routes

import (
	"net/http"
	model "ticketsystem/models"

	"github.com/gin-gonic/gin"
)

func getAllTickets(context *gin.Context) {
	tickets, err := model.GetTicketsFromDB()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "There was some internal server error", "err": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "This is a GET request", "tickets": tickets})
}

func createTicket(context *gin.Context) {
	var ticket model.Ticket
	err := context.ShouldBindJSON(&ticket)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "There was some error parsing the ticket", "err": err.Error()})
	}
	err = ticket.SaveToDB()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "There was some error", "err": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "New ticket created", "ticket": ticket})

}
