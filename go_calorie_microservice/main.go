package main

import (
	"fmt"
	"net/http"
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

func main() {
	r := gin.Default()
	r.GET("/caloric_goal/user/:user_id", getUserCaloricGoal)
	r.Run(":8080") // Listen and service on 0.0.0.0:8080/ping
}

func getUserCaloricGoal(c *gin.Context) {
	_id := c.Param("user_id")
	fmt.Println("id from query => ", _id)

	user_id, err := strconv.Atoi(_id)
	fmt.Println("ID => ", user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	goals, exists := findUserCalorieGoal(user_id)

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"Caloric goal does not exist for user with id:": user_id,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"goal_for_user": goals,
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
