package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

var ticketsList = tickets.Tickets{}

func main() {
	content, err := tickets.LoadTicketCSV("tickets.csv")
	if err != nil {
		fmt.Println(err)
	}

	tickets.DumpTicketData(content, &ticketsList)

	// Print all the ticket list
	// fmt.Println(ticketsList.GetListTickets())

	var destination = "Brazil"
	total, err := ticketsList.GetTotalTickets(destination)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Total tickets to %s: %d\n", destination, total)

	dayStage := "morning"
	totalCountStage, err := ticketsList.GetCountByPeriod(dayStage)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Total tickets in the %s: %d\n", dayStage, totalCountStage)

	totalCountAvg, err := ticketsList.AverageDestination(destination)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Average tickets to %s: %.2f%%\n", destination, totalCountAvg)

}
