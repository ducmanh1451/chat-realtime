package repositories

import (
	"chat-realtime/internal/helpers"
	"chat-realtime/internal/models"

	"gorm.io/gorm"
)

type FriendRepository struct {
	DB *gorm.DB
}

// init repository
func InitFriendRepository(db *gorm.DB) *FriendRepository {
	return &FriendRepository{DB: db}
}

// send friend request
func (r *FriendRepository) SendFriendRequest(sender_id int, sender_name string, receiver_id int) (*models.Friend, *helpers.Response) {
	var existingFriend models.Friend

	if sender_id == receiver_id {
		return nil, helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E001"))
	}

	if err := r.DB.Table("friend").Where("(sender_id = ? AND receiver_id = ? OR sender_id = ? AND receiver_id = ?) AND deleted_at IS NULL", sender_id, receiver_id, receiver_id, sender_id).First(&existingFriend).Error; err == nil {
		// if exists request => check status
		switch existingFriend.Status {
		case 0:
			return nil, helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E009"))
		case 1:
			return nil, helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E010"))
		}
	}

	tx := r.DB.Begin()
	if tx.Error != nil {
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	newFriend := models.Friend{
		SenderID:   sender_id,
		ReceiverID: receiver_id,
		Status:     0,
	}
	// insert table friend
	if err := r.DB.Table("friend").Omit("updated_at").Create(&newFriend).Error; err != nil {
		tx.Rollback()
		return nil, helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E001"))
	}
	newNotification := models.Notification{
		SenderID:   sender_id,
		ReceiverID: receiver_id,
		Message:    sender_name + " đã gửi lời mời kết bạn!",
	}
	// insert table notification
	if err := r.DB.Table("notification").Omit("updated_at").Create(&newNotification).Error; err != nil {
		tx.Rollback()
		return nil, helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E001"))
	}
	if err := tx.Commit().Error; err != nil {
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	return nil, helpers.CreateResponse(helpers.StatusOK, helpers.GetMessage("S001"))
}

// accept friend request
func (r *FriendRepository) AcceptFriendRequest(sender_id int, receiver_id int, receiver_name string) (*models.Friend, *helpers.Response) {
	var existingFriend models.Friend
	if err := r.DB.Table("friend").Where("sender_id = ? AND receiver_id = ? AND deleted_at IS NULL", sender_id, receiver_id).First(&existingFriend).Error; err != nil {
		return nil, helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E004"))
	}

	tx := r.DB.Begin()
	if tx.Error != nil {
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	// update table friend
	if err := r.DB.Table("friend").Model(&existingFriend).Updates(map[string]interface{}{
		"status":     1,
		"updated_at": gorm.Expr("CURRENT_TIMESTAMP"),
	}).Error; err != nil {
		tx.Rollback()
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	// insert table notification
	newNotification := models.Notification{
		SenderID:   receiver_id,
		ReceiverID: sender_id,
		Message:    receiver_name + " đã chấp nhận lời mời kết bạn!",
	}
	if err := r.DB.Table("notification").Omit("updated_at").Create(&newNotification).Error; err != nil {
		tx.Rollback()
		return nil, helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E001"))
	}

	if err := tx.Commit().Error; err != nil {
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	return nil, helpers.CreateResponse(helpers.StatusOK, helpers.GetMessage("S001"))
}

// delete friend request
func (r *FriendRepository) DeleteFriendRequest(sender_id int, receiver_id int) (*models.Friend, *helpers.Response) {
	var existingFriend models.Friend
	// check exist
	if err := r.DB.Table("friend").Where("(sender_id = ? AND receiver_id = ? OR sender_id = ? AND receiver_id = ?) AND deleted_at IS NULL", sender_id, receiver_id, receiver_id, sender_id).First(&existingFriend).Error; err != nil {
		return nil, helpers.CreateResponse(helpers.StatusBadRequest, helpers.GetMessage("E004"))
	}

	tx := r.DB.Begin()
	if tx.Error != nil {
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	// update table friend
	if err := r.DB.Table("friend").Model(&existingFriend).Updates(map[string]interface{}{
		"status":     3,
		"deleted_at": gorm.Expr("CURRENT_TIMESTAMP"),
	}).Error; err != nil {
		tx.Rollback()
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}
	if err := tx.Commit().Error; err != nil {
		return nil, helpers.CreateResponse(helpers.StatusInternalServerError, helpers.GetMessage("E001"))
	}

	return nil, helpers.CreateResponse(helpers.StatusOK, helpers.GetMessage("S001"))
}
