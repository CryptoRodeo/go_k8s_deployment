package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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
	fmt.Printf("Now running %s on port %s\n", getEnv("APP_NAME", "go-users"), getEnv("PORT", "8080"))

	r := gin.Default()
	r.GET("/api/v1/users", getUsers)
	r.GET("/api/v1/calorie_goals/user/:id", getUsersCalorieTargets)
	r.Run(fmt.Sprintf(":%s", getEnv("PORT", "8080")))
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
		"users": users,
	})
}

func getUsersCalorieTargets(c *gin.Context) {
	order_service_url := getEnv("ORDER", "localhost:8081")
	resp, err := http.Get(order_service_url)

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
