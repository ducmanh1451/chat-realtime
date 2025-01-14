package helpers

var ErrorMessages = map[string]string{
	"E001": "Lỗi hệ thống!",
	"E002": "Email đã tồn tại!",
	"E003": "Dữ liệu không hợp lệ!",
	"E004": "Dữ liệu không tồn tại!",
	"E005": "Tạo session token không thành công!",
	"E006": "Không tìm thấy session token!",
	"E007": "Xảy ra lỗi khi xóa session token!",
	"E008": "Session token không hợp lệ hoặc đã hết hạn!",
}

func GetErrorMessage(code string) string {
	if msg, exists := ErrorMessages[code]; exists {
		return msg
	}
	return ErrorMessages["E001"]
}
