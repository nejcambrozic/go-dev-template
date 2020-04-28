package health

type ApiUserMessage struct {
	Status string
}

type Service interface {
	GetHealth() (ApiUserMessage, int)
}

type service struct {
	name string
}

func NewService() Service {
	return &service{"basic"}
}

func (s *service) GetHealth() (ApiUserMessage, int) {
	healthMessage := ApiUserMessage{Status: "ok"}
	status := 200
	return healthMessage, status
}
