package user

import "log"

type Service interface { // Contrato
	Create(firsName, lastName, email, phone string) error
}

type service struct{} // Sera la Implementción de la interfaz

func NewService() Service { // Constructor -> interfaz
	return &service{}
}

func (s service) Create(firsName, lastName, email, phone string) error {
	log.Println("Create user service")
	return nil
}
