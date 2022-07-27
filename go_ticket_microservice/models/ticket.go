package models

type Ticket struct {
	Id          int    `json:"ticket_id"`
	Description string `json:"description"`
	IsUrgent    bool   `json:"is_urgent"`
}

func ExistsByID(id int) bool {
	for _, ticket := range fakeTicketDB() {
		if ticket.Id == id {
			return true
		}
	}
	return false
}

func GetTicket(id int) *Ticket {
	for _, ticket := range fakeTicketDB() {
		if ticket.Id == id {
			return &ticket
		}
	}
	return nil
}

func GetTicketsByID(ids []int) []*Ticket {
	temp := []*Ticket{}
	for _, ticket := range fakeTicketDB() {
		if inIdList(ids, ticket.Id) {
			temp = append(temp, &ticket)
		}
	}
	return temp
}

func GetAllTickets() []Ticket {
	return fakeTicketDB()
}

func inIdList(list []int, val int) bool {
	for _, list_val := range list {
		if list_val == val {
			return true
		}
	}
	return false
}

func fakeTicketDB() []Ticket {
	return []Ticket{
		{1800, "The netflix wont work!", true},
		{1500, "Need new mouse", false},
		{2000, "Donut ordering microservice is down!", true},
		{2001, "Pizza ordering microservice is down too!", true},
	}
}
