package health

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) Check() HealthResponse {
	return HealthResponse{
		Status: "OK",
	}
}
