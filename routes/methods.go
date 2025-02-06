package routes

import (
	"net/http"
	"strconv"
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

func getSingleTicket(context *gin.Context) {
	ticketID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ticket id", "err": err.Error()})
		return
	}
	ticket, err := model.GetTicketByID(ticketID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "There was some error", "err": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Ticket details", "ticket": ticket})
}

func updateTicket(context *gin.Context) {
	ticketID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cold not find ticket with the id", "err": err.Error()})
		return
	}
	_, err = model.GetTicketByID(ticketID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "There was some interal server error", "err": err.Error()})
		return
	}
	var updatedTicket model.Ticket
	err = context.ShouldBindJSON(&updatedTicket)
	updatedTicket.ID = ticketID
	err = updatedTicket.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "There was some internal error", "err": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Ticket updated successfully", "ticket": updatedTicket})

}
