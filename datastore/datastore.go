package datastore

import (
	"fmt"
	"sync"
	"testapi/models"
)

type ContactBook interface {
	AddContact(contact *models.Contact) error
	GetContacts() ([]models.Contact, error)
	DeleteContact(contactName string) error
	UpdateContact(*models.Contact) error

}

type InMemory struct {
	*sync.Mutex
	Contacts map[string]string
}

func (m *InMemory) AddContact(contact *models.Contact) error {
	m.Lock()
	m.Contacts[contact.Name] = contact.Phone
	m.Unlock()
	return nil
}

func (m *InMemory) GetContacts() ([]models.Contact, error) {
	m.Lock()
	cts := make([]models.Contact, 0, len(m.Contacts))

	for ctK, ctV := range m.Contacts {
		cts = append(cts, models.Contact{Name: ctK, Phone: ctV})
	}
	m.Unlock()

	return cts, nil
}

func (m *InMemory) DeleteContact(contactNumber string) error {

	for key, val := range m.Contacts {
		if val == contactNumber {
			delete(m.Contacts, key)
			return nil
		}
	}
	return fmt.Errorf("%v Not Found", contactNumber)
	
}

func (m *InMemory) UpdateContact(contact *models.Contact) error {
	if _, ok := m.Contacts[contact.Name]; ok {
		m.Contacts[contact.Name] = contact.Phone
		return nil
	} 
	m.AddContact(contact)
	return fmt.Errorf("contact not found, so added")
}