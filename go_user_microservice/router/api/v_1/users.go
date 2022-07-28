package v_1

import (
	"encoding/json"
	"fmt"
	"go_user_microservice/service/user_service"
	"go_user_microservice/service/util/settings"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Params struct {
	UserIds   []int `json:"ids"`
	TicketIds []int `json:"ticket_ids"`
}

func GetAllUsers(c *gin.Context) {
	users := user_service.GetAllUSers()

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUser(c *gin.Context) {
	_id := c.Param("id")

	id, err := strconv.Atoi(_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	if !user_service.ExistsByID(id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("User at id %d does not exist.", id),
		})
		return
	}

	user := user_service.GetUserByID(id)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func SearchUsers(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	//request_body := string(jsonData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	params := Params{}
	json.Unmarshal([]byte(jsonData), &params)
	fmt.Println(params.UserIds)

	usersFound := user_service.SearchUsers(params.UserIds)

	c.JSON(http.StatusOK, gin.H{
		"data": usersFound,
	})
}

func GetUsersTickets(c *gin.Context) {
	endpoint := (settings.ServiceEndpoints["ticket-service"] + "/search")
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	user := user_service.GetUserByID(id)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("User at ID: ", id, " does not exists."),
		})
	}

	resp, err := http.NewRequest("GET", endpoint, user.Tickets)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var result map[string]interface{}
	body_text := string(body)
	json.Unmarshal([]byte(body_text), &result)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
