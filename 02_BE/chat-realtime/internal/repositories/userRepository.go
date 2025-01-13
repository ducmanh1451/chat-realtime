package repositories

import (
	"chat-realtime/internal/helpers"
	"chat-realtime/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// init repository
func InitUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// get users
func (r *UserRepository) GetUsers() ([]models.User, *helpers.Error) {
	var users []models.User
	if err := r.DB.Table("user").Find(&users).Error; err != nil {
		return nil, helpers.CreateError(helpers.StatusInternalServerError, helpers.GetErrorMessage("E001"))
	}
	return users, nil
}

// create user
func (r *UserRepository) CreateUser(user *models.User) (*models.User, *helpers.Error) {
	var existingUser models.User
	if err := r.DB.Table("user").Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return nil, helpers.CreateError(helpers.StatusBadRequest, helpers.GetErrorMessage("E002"))
	}

	if err := r.DB.Table("user").Create(user).Error; err != nil {
		return nil, helpers.CreateError(helpers.StatusInternalServerError, helpers.GetErrorMessage("E001"))
	}
	return user, nil
}
