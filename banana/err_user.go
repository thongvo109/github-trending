package banana

import "errors"

var (
	UserConflict = errors.New("Người dùng đã tồn tại")
	UserNotFound = errors.New("Người dùng không tồn tại")
	SignUpFail   = errors.New("Đang kí thất bại")
	UserNotUpdated = errors.New("Cập nhật thông tin người dùng thất bại")
)
