package v_1

import (
	"encoding/json"
	"go_ticket_microservice/service/ticket_service"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Params struct {
	TicketIds []int `json:"tickets" binding:"required"`
}

func GetAllTickets(c *gin.Context) {
	tickets := ticket_service.GetAllTickets()

	c.JSON(http.StatusOK, gin.H{
		"data": tickets,
	})
}

func SearchTickets(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	request_body := string(jsonData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	params := Params{}
	json.Unmarshal([]byte(request_body), &params)
	ticketsFound := ticket_service.SearchTickets(params.TicketIds)

	c.JSON(http.StatusOK, gin.H{
		"data": ticketsFound,
	})
}
