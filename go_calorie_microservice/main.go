package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserCalories struct {
	UserId       int `json:"user_id"`
	Caloric_Goal int `json:"caloric_goal"`
}

// For now we'll just throw this dummy data out.
var calories = []UserCalories{
	{1, 1800},
	{2, 1500},
	{3, 20000},
}
var PORT string = fmt.Sprintf(":%s", getEnv("PORT", "8081"))
var APP_NAME string = getEnv("APP_NAME", "go-calorie-calc")

func main() {
	fmt.Printf("Now running %s on port %s\n", APP_NAME, PORT)

	r := gin.Default()
	r.GET("/api/v1/calorie_goal/user/:user_id", getUserCaloricGoal)
	r.Run(PORT)
}

func getUserCaloricGoal(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("user_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	goal, exists := findUserCalorieGoal(user_id)

	if !exists {
		error_msg := fmt.Sprintf("Caloric goal does not exist for user with ID %d", user_id)

		c.JSON(http.StatusBadRequest, gin.H{
			"Error": error_msg,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"calorie_goal": goal,
	})
}

func findUserCalorieGoal(user_id int) (goal int, exists bool) {
	for _, uc := range calories {
		if uc.UserId == user_id {
			return uc.Caloric_Goal, true
		}
	}
	return 0, false
}

func getEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return fallback
}
