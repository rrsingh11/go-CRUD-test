package services


type ValidationService struct {
	noOfDigits int
}


func NewValidationSevice(d int) *ValidationService{
	return &ValidationService{
		noOfDigits: d,
	}
}

func (s *ValidationService) Check(phoneNumber string) bool {
	return len(phoneNumber) == s.noOfDigits
}