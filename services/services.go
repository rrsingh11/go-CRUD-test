package services

import "testapi/models"


type Services struct {
	Contactservice models.Service
}


func NewServices(cs models.Service) *Services {
	return &Services{
		Contactservice: cs,
	}
}
