package user_service

import (
	"go_user_microservice/models"
)

func GetAllUsers() []models.User {
	return models.GetAllUsers()
}

func GetUserByID(id int) *models.User {
	return models.GetUser(id)
}

func SearchUsers(ids []int) []models.User {
	return models.GetUsersByID(ids)
}

func ExistsByID(id int) bool {
	return models.ExistsByID(id)
}
