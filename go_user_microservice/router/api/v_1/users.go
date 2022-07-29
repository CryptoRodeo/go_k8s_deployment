package v_1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_user_microservice/service/user_service"
	"go_user_microservice/utils/settings"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserParams struct {
	Ids []int `json:"ids" binding:"required"`
}

func GetAllUsers(c *gin.Context) {
	users := user_service.GetAllUsers()

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUserByID(c *gin.Context) {
	_id := c.Param("id")

	id, err := strconv.Atoi(_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid 'id' value passed.",
		})
		return
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

	params := UserParams{}
	json.Unmarshal([]byte(jsonData), &params)

	usersFound := user_service.SearchUsers(params.Ids)

	c.JSON(http.StatusOK, gin.H{
		"data": usersFound,
	})
}

type TicketParams struct {
	TicketIds []int `json:"ticket_ids"`
}

func GetUsersTickets(c *gin.Context) {
	endpoint := (settings.ServiceEndpoints["ticket-service"] + "/search")
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user := user_service.GetUserByID(id)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("User at ID: %d does not exists.", id),
		})
		return
	}

	jsonData := map[string][]int{
		"ticket_ids": user.Tickets,
	}

	req_body, _ := json.Marshal(jsonData)
	resp, err := MakeRequest(http.MethodGet, endpoint, bytes.NewBuffer(req_body))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	/**
	notes on resp.body:
		- don't close the body until we return from the function.
		- resp.Body is a stream of data read by the http client.
		- Do not forget to add this closing instruction,
		  otherwise, the client might not reuse a potential persistent connection to the server
	**/
	defer resp.Body.Close()
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

	c.JSON(http.StatusOK, result)
}

// Wrapper for Http#NewRequest
func MakeRequest(method string, url string, body io.Reader) (*http.Response, error) {
	// Create new client, build the request.
	cl := &http.Client{}
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	// add header, make request.
	req.Header.Add("Accept", `application/json`)
	resp, err := cl.Do(req)

	return resp, err
}
