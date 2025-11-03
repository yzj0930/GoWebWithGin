package services

import (
	"github.com/yzj0930/GoWebWithGin/dto/request"
	"github.com/yzj0930/GoWebWithGin/dto/response"
)

type PingService struct {
}

func (s *PingService) GetPingMessage() string {
	return "pong"
}

func (s *PingService) GetPostJson(request request.PostJsonRequest) response.PostJsonResponse {
	return response.PostJsonResponse{
		Name:  request.Name,
		Email: request.Email,
		Age:   request.Age,
	}
}
