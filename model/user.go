package model

// Tạo struct để lưu thông tin của một bản ghi trong bảng "users"
type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRequest struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}
