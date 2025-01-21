package helpers

var Messages = map[string]string{
	// errors
	"E001": "Lỗi hệ thống!",
	"E002": "Email đã tồn tại!",
	"E003": "Dữ liệu không hợp lệ!",
	"E004": "Dữ liệu không tồn tại!",
	"E005": "Tạo session token không thành công!",
	"E006": "Không tìm thấy session token!",
	"E007": "Xảy ra lỗi khi xóa session token!",
	"E008": "Session token không hợp lệ hoặc đã hết hạn!",
	"E009": "Lời mời kết bạn đã được tạo!",
	"E010": "Hai người đã là bạn bè!",
	// success
	"S001": "Thành công!",
}

func GetMessage(code string) string {
	if msg, exists := Messages[code]; exists {
		return msg
	}
	return Messages["E001"]
}
