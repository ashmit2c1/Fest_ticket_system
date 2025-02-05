package model

import (
	"fmt"
	"ticketsystem/db"
)

type Ticket struct {
	ID                  int64
	Name                string `binding:"required"`
	Registration_number string `binding:"required"`
	Day1                string `binding:"required"`
	Day2                string `binding:"required"`
	Day3                string `binding:"required"`
}

/*var tickets = []Ticket{}

func (ticket Ticket) Save() {
	tickets = append(tickets, ticket)
}

func GetAllTickets() []Ticket {
	return tickets
}*/

func (ticket Ticket) SaveToDB() error {
	query := `INSERT INTO tickets(name, registration_number, day1, day2, day3) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(ticket.Name, ticket.Registration_number, ticket.Day1, ticket.Day2, ticket.Day3)
	if err != nil {
		return err
	}
	ticket_id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	ticket.ID = ticket_id
	fmt.Println("Ticket saved to the database")
	return nil

}

func GetTicketsFromDB() ([]Ticket, error) {
	query := `SELECT * FROM tickets`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tickets []Ticket
	for rows.Next() {
		var ticket Ticket
		err := rows.Scan(&ticket.ID, &ticket.Name, &ticket.Registration_number, &ticket.Day1, &ticket.Day2, &ticket.Day3)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}
