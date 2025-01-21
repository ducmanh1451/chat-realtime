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
func (r *UserRepository) GetUsers() ([]models.User, *helpers.Response) {
	var users []models.User
	if err := r.DB.Table("user").Find(&users).Error; err != nil {
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	return users, helpers.CreateResponse(helpers.StatusOK, helpers.GetMessage("S001"))
}

// create user
func (r *UserRepository) CreateUser(user *models.User) (*models.User, *helpers.Response) {
	var existingUser models.User
	r.DB.Table("user").Where("email = ? AND deleted_at IS NULL", user.Email).First(&existingUser)
	if existingUser.ID != 0 {
		return nil, helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E002"))
	}

	tx := r.DB.Begin()
	if tx.Error != nil {
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	if err := r.DB.Table("user").Create(user).Error; err != nil {
		tx.Rollback()
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	if err := tx.Commit().Error; err != nil {
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	return user, helpers.CreateResponse(helpers.StatusOK, helpers.GetMessage("S001"))
}

// update user
func (r *UserRepository) UpdateUser(id string, userUpdate *models.User) (*models.User, *helpers.Response) {
	var existingUser models.User

	if err := r.DB.Table("user").Where("id = ?", id).First(&existingUser).Error; err != nil {
		return nil, helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E004"))
	}

	tx := r.DB.Begin()
	if tx.Error != nil {
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	if err := r.DB.Table("user").Model(&existingUser).Updates(map[string]interface{}{
		"password":      userUpdate.Password,
		"full_name":     userUpdate.FullName,
		"gender":        userUpdate.Gender,
		"date_of_birth": userUpdate.DateOfBirth,
		"updated_at":    gorm.Expr("CURRENT_TIMESTAMP"),
	}).Error; err != nil {
		tx.Rollback()
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	if err := tx.Commit().Error; err != nil {
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}

	var updatedUser models.User
	if err := r.DB.Table("user").Where("id = ?", id).First(&updatedUser).Error; err != nil {
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}

	return &updatedUser, helpers.CreateResponse(helpers.StatusOK, helpers.GetMessage("S001"))
}
