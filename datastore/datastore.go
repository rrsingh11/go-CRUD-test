package datastore

import "testapi/models"

type ContactBook interface {
	AddContact(contact *models.Contact) error
	GetContacts() ([]models.Contact, error)
	DeleteContact( string) error
	UpdateContact(*models.Contact) error
	InsertManyContacts([]models.Contact) error
}