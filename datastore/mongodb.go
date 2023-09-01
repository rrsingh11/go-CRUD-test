package datastore

import "testapi/models"


type MongoDB struct {

}

func (m *MongoDB) AddContact(contact *models.Contact) error {

	return nil
}

func (m *MongoDB) GetContacts() ([]models.Contact, error) {

	return nil, nil
}

func (m *MongoDB) DeleteContact(contactNumber string) error {

	return nil
}

func (m *MongoDB) UpdateContact(contact *models.Contact) error {
	
	return nil
}