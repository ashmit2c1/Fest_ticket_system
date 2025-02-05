package model

type Ticket struct {
	ID                  int64
	Name                string `binding:"required"`
	Registration_number string `binding:"required"`
	Day1                string `binding:"required"`
	Day2                string `binding:"required"`
	Day3                string `binding:"required"`
}

var tickets = []Ticket{}

func (ticket Ticket) Save() {
	tickets = append(tickets, ticket)
}

func GetAllTickets() []Ticket {
	return tickets
}
