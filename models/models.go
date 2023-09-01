package models
// ContactInfo
type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Service interface {
	Check() bool
}

type Output interface {
	Decodable()
}

type ContactBook interface {
	AddContact(contact *Contact) error
	GetContacts() (Output, error)
	DeleteContact(contactName string) error
	UpdateContact(*Contact) error

}


