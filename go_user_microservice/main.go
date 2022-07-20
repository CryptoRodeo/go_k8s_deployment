package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"data"

	"github.com/gin-gonic/gin"
)

var PORT string = fmt.Sprintf(":%s", getEnv("PORT", "8080"))
var APP_NAME string = getEnv("APP_NAME", "go-users")
var TICKET_SERVICE = getEnv("TICKET_SERVICE", "localhost:8081")

func main() {
	fmt.Printf("Now running %s on port %s\n", APP_NAME, PORT)

	r := gin.Default()
	r.GET("/api/v1/users", getUsers)
	r.GET("/api/v1/user/:id/tickets", getUsersTickets)
	r.Run(PORT)
}

func getEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return fallback
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": data.USERS,
	})
}

func getUsersTickets(c *gin.Context) {
	user_id := c.Param("id")
	endpoint := fmt.Sprintf("http://%s/api/v1/tickets/user/%s", TICKET_SERVICE, user_id)

	resp, err := http.Get(endpoint)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var result map[string]interface{}
	body_text := string(body)
	json.Unmarshal([]byte(body_text), &result)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
