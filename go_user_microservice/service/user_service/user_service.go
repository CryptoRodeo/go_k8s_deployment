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

func ExistsByID(id int) bool {
	return models.ExistsByID(id)
}
