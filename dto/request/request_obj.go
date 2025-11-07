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
