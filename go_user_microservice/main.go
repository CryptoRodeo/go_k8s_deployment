package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	CurrentWeight int    `json:"current_weight"`
	GoalWeight    int    `json:"goal_weight"`
}

var users = map[string]User{
	"joey": {
		Id:            1,
		Name:          "Joey",
		CurrentWeight: 180,
		GoalWeight:    160,
	},
	"sandy": {
		Id:            2,
		Name:          "Sandy",
		CurrentWeight: 150,
		GoalWeight:    130,
	},
	"anthony": {
		Id:            3,
		Name:          "Anthony",
		CurrentWeight: 200,
		GoalWeight:    180,
	},
}

func main() {
	r := gin.Default()
	r.GET("/users", getUsers)
	r.GET("/user/:id/calorie_targets", getUsersCalorieTargets)
	r.Run(":8081") // Listen and service on 0.0.0.0:8080/ping
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

/**
WIP, need to set up:
- environment variables
- other microservice
**/
func getUsersCalorieTargets(c *gin.Context) {
	// Get sample data from this endpoint
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

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
		"calorie_targets": result,
	})
}
