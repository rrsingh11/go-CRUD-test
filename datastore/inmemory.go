package datastore

import (
	"fmt"
	"sync"
	"testapi/models"
)

type InMemory struct {
	mu sync.Mutex
	Contacts map[string]string
}

func (m *InMemory) AddContact(contact *models.Contact) error {
	m.mu.Lock()
	m.Contacts[contact.Name] = contact.Phone
	m.mu.Unlock()
	return nil
}

type mapData []models.Contact
func (md mapData) Decodable() {}

func (m *InMemory) GetContacts() (models.Output, error) {
	m.mu.Lock()
	cts := make(mapData, 0, len(m.Contacts))

	for ctK, ctV := range m.Contacts {
		cts = append(cts, models.Contact{Name: ctK, Phone: ctV})
	}
	m.mu.Unlock()

	return cts, nil
}

func (m *InMemory) DeleteContact(contactNumber string) error {
	m.mu.Lock()
	for key, val := range m.Contacts {
		if val == contactNumber {
			delete(m.Contacts, key)
			m.mu.Unlock()
			return nil
		}
	}
	m.mu.Unlock()
	return fmt.Errorf("%v Not Found", contactNumber)
	
}

func (m *InMemory) UpdateContact(contact *models.Contact) error {
	m.mu.Lock()
	if _, ok := m.Contacts[contact.Name]; ok {
		m.Contacts[contact.Name] = contact.Phone
		m.mu.Unlock()
		return nil
	} 
	m.mu.Unlock()
	return fmt.Errorf("contact not found")
}