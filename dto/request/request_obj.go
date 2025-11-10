package request

type PostJsonRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type UserRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type UserListRequest struct {
	// 可以根据需要添加过滤条件，例如分页参数等
	Name     string `json:"name"`
	Code     string `json:"code"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}
