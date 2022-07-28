package ticket_service

import (
	"go_ticket_microservice/models"
)

func GetAllTickets() []models.Ticket {
	return models.GetAllTickets()
}

func GetTicketByID(id int) *models.Ticket {
	return models.GetTicket(id)
}

func SearchTickets(ids []int) []models.Ticket {
	return models.GetTicketsByID(ids)
}

func ExistsByID(id int) bool {
	return models.ExistsByID(id)
}
