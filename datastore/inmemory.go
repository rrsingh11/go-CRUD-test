package datastore

import (
	"sync"
	"testapi/models"
)

type InMemory struct {
	mu sync.Mutex
	Contacts map[string]string
}

func (m *InMemory) AddContact(contact *models.Contact) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Contacts[contact.Name] = contact.Phone
	return nil
}



func (m *InMemory) GetContacts() ([]models.Contact, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	cts := make([]models.Contact, 0, len(m.Contacts))

	for ctK, ctV := range m.Contacts {
		cts = append(cts, models.Contact{Name: ctK, Phone: ctV})
	}
	if len(cts) == 0 {
		return cts, models.ErrNotFound
	}
	return cts, nil
}

func (m *InMemory) DeleteContact(contactNumber string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for key, val := range m.Contacts {
		if val == contactNumber {
			delete(m.Contacts, key)
			return nil
		}
	}
	return models.ErrNotFound
	
}

func (m *InMemory) UpdateContact(contact *models.Contact) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.Contacts[contact.Name]; ok {
		m.Contacts[contact.Name] = contact.Phone
		return nil
	} 
	
	return models.ErrNotFound
}

func (m *InMemory) InsertManyContacts(contacts []models.Contact) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	for _, contact := range contacts {
		m.Contacts[contact.Name] = contact.Phone
	}
	return nil
}