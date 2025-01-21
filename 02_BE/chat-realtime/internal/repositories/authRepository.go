package repositories

import (
	"chat-realtime/internal/helpers"
	"chat-realtime/internal/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

// init repository
func InitAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

// login
func (r *AuthRepository) Login(email string, password string) (*models.User, *helpers.Response) {
	var user models.User
	// check exist user
	if err := r.DB.Table("user").Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil, helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E004"))
	}
	return &user, helpers.CreateResponse(helpers.StatusOK, helpers.GetMessage("S001"))
}
