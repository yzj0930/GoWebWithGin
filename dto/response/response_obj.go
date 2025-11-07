package response

import "time"

type ResponseDto struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PostJsonResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type UserResponseDto struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	CreatedTime time.Time `json:"create_time"`
	UpdatedTime time.Time `json:"update_time"`
	// DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`   // 软删除字段暂时不考虑，后续需要再加
}
