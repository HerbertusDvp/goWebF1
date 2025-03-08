package user

import "log"

type Service interface {
	Create(firsName, lastName, email, phone string) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s service) Create(firsName, lastName, email, phone string) error {
	log.Println("Create user service")
	return nil
}
