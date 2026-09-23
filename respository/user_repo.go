package respository

import (
	"strconv"

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

func DeleteUser(id string) (bool, error) {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return false, err
	}
	result := config.DB.Delete(&models.User{}, uid)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
