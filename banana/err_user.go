package banana

import "errors"

var (
	UserConflict = errors.New("Người dùng đã tồn tại")
	SignUpFail   = errors.New("Đăng ký thất bại")
	UserNotFound = errors.New("Người dùng không tồn tại")
)
