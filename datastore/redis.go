package datastore

import (
	"testapi/models"

	"github.com/go-redis/redis"
)

type Redis struct {
	Client redis.Client
}

func (r Redis) AddContact(contact *models.Contact) error {
	err := r.Client.HSetNX("contacts", contact.Name, contact.Phone).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r Redis) GetContacts() ([]models.Contact, error) {
	res, err := r.Client.HGetAll("contacts").Result()
	if err != nil {
		return nil, err
	}

	var contacts []models.Contact

	for k, v := range res {
		contacts = append(contacts, models.Contact{Name: k, Phone: v})
	}

	if len(contacts) == 0 {
		return contacts, models.ErrNotFound
	}
	return contacts, nil
}

func (r Redis) DeleteContact(contactNumber string) error {
	// db call
	cts, nil := r.GetContacts()
	var key string
	// linear search the contact 
	for _, ct := range cts{
		if ct.Phone == contactNumber {
			key = ct.Name
			break
		}
	}

	err := r.Client.HDel("contacts", key).Err()

	if err != nil {
		return err
	}



	return nil
}

func (r Redis) UpdateContact(contact *models.Contact) error {
	err := r.Client.HSet("contacts", contact.Name, contact.Phone).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r Redis) InsertManyContacts(contacts []models.Contact) error {
	cts := make(map[string]interface{})

	for _, contact := range contacts {
		cts[contact.Name] = contact.Phone
	}

	err := r.Client.HMSet("contacts", cts).Err()

	if err != nil {
		return err
	}

	return nil
}
