package services

import (
	"menu-delivery/config"
	"menu-delivery/models"
)

type UserInputDetails struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Address  *string `json:"address"`
	Phone    *string `json:"phone"`
	CPF      *string `json:"cpf"`
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id string) (models.User, error) {
	var user models.User

	if error := config.DB.Where("id = ?", id).First(&user).Error; error != nil {
		return models.User{}, error
	}
	return user, nil
}

func CreateUser(input UserInputDetails) (*models.User, error) {
	user := models.User{}

	if input.Name != nil {
		user.Name = *input.Name
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	if input.Address != nil {
		user.Address = *input.Address
	}

	if input.Password != nil {
		user.Password = *input.Password
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id string, input UserInputDetails) (*models.User, error) {
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return &models.User{}, nil
	}

	if input.Name != nil {
		user.Name = *input.Name
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	if input.Address != nil {
		user.Address = *input.Address
	}

	if input.Password != nil {
		user.Password = *input.Password
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return &models.User{}, nil
	}

	return &user, nil

}

func DeleteUser(id string) (error, bool) {
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return err, false
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return err, false
	}

	return nil, true
}
