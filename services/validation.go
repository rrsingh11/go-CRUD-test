package services

import "testapi/models"


type IsValidPhoneNumber struct {
	phoneNumber string
}


func NewValidationSevice(ph string) models.Service{
	return &IsValidPhoneNumber{
		phoneNumber: ph,
	}
}

func (s *IsValidPhoneNumber) Check() bool {
	return len(s.phoneNumber) == 10
}