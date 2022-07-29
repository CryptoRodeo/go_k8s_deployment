package v_1

import (
	"encoding/json"
	"fmt"
	"go_ticket_microservice/service/ticket_service"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Params struct {
	TicketIds []int `json:"ticket_ids" binding:"required"`
}

func GetAllTickets(c *gin.Context) {
	tickets := ticket_service.GetAllTickets()

	c.JSON(http.StatusOK, gin.H{
		"data": tickets,
	})
}

func GetTicket(c *gin.Context) {
	_id := c.Param("id")

	id, err := strconv.Atoi(_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if !ticket_service.ExistsByID(id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Ticket at id %d does not exist.", id),
		})
		return
	}

	ticket := ticket_service.GetTicketByID(id)
	c.JSON(http.StatusOK, gin.H{
		"data": ticket,
	})
}

func SearchTickets(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	params := Params{}
	json.Unmarshal(jsonData, &params)

	ticketsFound := ticket_service.SearchTickets(params.TicketIds)

	c.JSON(http.StatusOK, gin.H{
		"data": ticketsFound,
	})
}
