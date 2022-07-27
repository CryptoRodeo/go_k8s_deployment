package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"go_users_microservice/data"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var PORT string = fmt.Sprintf(":%s", getEnv("PORT", "8080"))
var APP_NAME string = getEnv("APP_NAME", "go-users")
var TICKET_SERVICE = getEnv("TICKET_SERVICE", "localhost:8081")

func main() {
	fmt.Printf("Now running %s on port %s\n", APP_NAME, PORT)

	r := gin.Default()

	// Allow requests to come from any origin.
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

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

	user, exists := getUserByID(user_id)

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("User at ID: ", user_id, " does not exists."),
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

func getUserByID(user_id int) (data.User, bool) {
	for _, user := range data.USERS {
		if user.Id == user_id {
			return user, true
		}
	}
	return nil, false
}
