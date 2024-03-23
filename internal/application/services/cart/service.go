package cart

import (
	"go-skeleton/internal/application/services"
)

type Service struct {
	services.BaseService
	response *Response
}

func NewService(log services.Logger) *Service {
	return &Service{
		BaseService: services.BaseService{
			Logger: log,
		},
	}
}

func (s *Service) Execute(request Request) {

	if err := request.Validate(); err != nil {
		s.BadRequest(request, err)
		return
	}

	s.produceResponseRule(request.Data)
}

func (s *Service) GetResponse() (*Response, *services.Error) {
	return s.response, s.Error
}

func (s *Service) produceResponseRule(data *Data) {
	s.response = &Response{
		Data: "",
	}
}
