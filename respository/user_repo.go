package respository

import (
	"github.com/AKSHAYK0UL/graph-ql/config"
	"github.com/AKSHAYK0UL/graph-ql/models"
)

func GetUser() ([]models.User, error) {
	var Users []models.User
	err := config.DB.Find(&Users).Error
	return Users, err
}

func CreateUser(user *models.User) error {
	return config.DB.Create(&user).Error
}
