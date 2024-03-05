package tickets

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

func LoadTicketCSV(fileName string) (fileContent [][]string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	fileContent, err = reader.ReadAll()

	if err != nil {
		return
	}

	return
}

func DumpTicketData(data [][]string, tickets *Tickets) {
	for _, row := range data {
		id, _ := strconv.Atoi(row[0])
		name := row[1]
		email := row[2]
		destinationCountry := row[3]
		flightTime := row[4]
		price, _ := strconv.ParseFloat(row[5], 64)

		ticket := Ticket{ID: id, Name: name, Email: email, DestinationCountry: destinationCountry, FlightTime: flightTime, Price: price}

		tickets.AddTicket(ticket)
	}
}

func GetHourTime(flightTime string, sep string) (int, error) {
	var hours []string = strings.Split(flightTime, sep)
	hour, err := strconv.Atoi(hours[0])
	if err != nil {
		return 0, err
	}

	return hour, nil
}
