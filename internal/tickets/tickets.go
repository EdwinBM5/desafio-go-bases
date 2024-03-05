package tickets

import (
	"errors"
)

type Ticket struct {
	ID                 int
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         string
	Price              float64
}

type Tickets struct {
	tickets []Ticket
}

func (t *Tickets) GetListTickets() []Ticket {
	return t.tickets
}

func (t *Tickets) AddTicket(ticket Ticket) {
	t.tickets = append(t.tickets, ticket)
}

const (
	ErrInvalidTime = "Error: invalid time"
)

// ejemplo 1
func (t *Tickets) GetTotalTickets(destination string) (int, error) {
	var totalCount int
	for _, ticket := range t.tickets {
		if ticket.DestinationCountry == destination {
			totalCount++
		}
	}

	return totalCount, nil
}

// ejemplo 2
func (t *Tickets) GetCountByPeriod(time string) (int, error) {
	dayStages := map[string]int{
		"earlyMorning": 0,
		"mornings":     0,
		"afternoon":    0,
		"night":        0,
	}

	if time != "earlyMorning" && time != "morning" && time != "afternoon" && time != "night" {
		return 0, errors.New(ErrInvalidTime)
	}

	for _, ticket := range t.tickets {
		hour, err := GetHourTime(ticket.FlightTime, ":")
		if err != nil {
			return 0, err
		}

		switch {
		case hour >= 0 && hour <= 6:
			dayStages["earlyMorning"]++
		case hour >= 7 && hour <= 12:
			dayStages["morning"]++
		case hour >= 13 && hour <= 19:
			dayStages["afternoon"]++
		case hour >= 20 && hour <= 23:
			dayStages["night"]++
		}
	}

	return dayStages[time], nil
}

// ejemplo 3
func (t *Tickets) AverageDestination(destination string) (float32, error) {
	totalCount, err := t.GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}

	result := float32(totalCount) / float32(len(t.tickets)) * 100

	return result, nil
}
