package models
// ContactInfo
type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Service interface {
	Check() bool
}

type ContactBook interface {
	AddContact(contact *Contact) error
	GetContacts() ([]Contact, error)
	DeleteContact(contactName string) error
	UpdateContact(*Contact) error

}


