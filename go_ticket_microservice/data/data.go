package data

type Ticket struct {
	UserId      int    `json:"user_id"`
	TicketID    int    `json:"ticket_id"`
	Description string `json:"description"`
	IsUrgent    bool   `json:"is_urgent"`
}

var Tickets = []Ticket{
	{1, 1800, "The netflix wont work!", true},
	{2, 1500, "Need new mouse", false},
	{3, 2000, "Donut ordering microservice is down!", true},
	{3, 2001, "Pizza ordering microservice is down too!", true},
}
