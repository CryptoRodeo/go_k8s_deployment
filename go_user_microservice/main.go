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

var PORT string = fmt.Sprintf(":%s", getEnv("PORT", "8080"))
var APP_NAME string = getEnv("APP_NAME", "go-users")
var CALORIE_SERVICE = getEnv("CALORIE_SERVICE", "localhost:8081")

func main() {
	fmt.Printf("Now running %s on port %s\n", APP_NAME, PORT)

	r := gin.Default()
	r.GET("/api/v1/users", getUsers)
	r.GET("/api/v1/calorie_goals/user/:id", getUsersCalorieTargets)
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
		"users": users,
	})
}

func getUsersCalorieTargets(c *gin.Context) {
	user_id := c.Param("id")
	endpoint := fmt.Sprintf("http://%s/api/v1/calorie_goal/user/%s", CALORIE_SERVICE, user_id)

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
