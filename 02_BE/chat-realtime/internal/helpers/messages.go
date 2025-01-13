package helpers

var ErrorMessages = map[string]string{
	"E001": "Lỗi hệ thống!",
	"E002": "Email đã tồn tại!",
	"E003": "Dữ liệu không hợp lệ!",
}

func GetErrorMessage(code string) string {
	if msg, exists := ErrorMessages[code]; exists {
		return msg
	}
	return ErrorMessages["E001"]
}
