package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Ticket struct {
	UserId      int    `json:"user_id"`
	TicketID    int    `json:"ticket_id"`
	Description string `json:"description"`
	IsUrgent    bool   `json:"is_urgent"`
}

// For now we'll just throw this dummy data out.
var tickets = []Ticket{
	{1, 1800, "The netflix wont work!", true},
	{2, 1500, "Need new mouse", false},
	{3, 2000, "Donut ordering microservice is down!", true},
	{3, 2001, "Pizza ordering microservice is down too!", true},
}
var PORT string = fmt.Sprintf(":%s", getEnv("PORT", "8081"))
var APP_NAME string = getEnv("APP_NAME", "go-ticket")

func main() {
	fmt.Printf("Now running %s on port %s\n", APP_NAME, PORT)

	r := gin.Default()
	r.GET("/api/v1/tickets", getTickets)
	r.GET("/api/v1/tickets/user/:user_id", getUserTickets)
	r.Run(PORT)
}

func getTickets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": tickets,
	})
}

func getUserTickets(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	tickets, empty := findUserTickets(user_id)

	if empty {
		error_msg := fmt.Sprintf("No tickets found for user with id: %d", user_id)

		c.JSON(http.StatusBadRequest, gin.H{
			"Error": error_msg,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tickets,
	})
}

func findUserTickets(user_id int) ([]Ticket, bool) {
	temp := []Ticket{}
	for _, ticket := range tickets {
		if ticket.UserId == user_id {
			temp = append(temp, ticket)
		}
	}
	return temp, len(temp) == 0
}

func getEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return fallback
}
