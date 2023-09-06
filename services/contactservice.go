package services

import (
	"fmt"
	"io"
	"testapi/models"
	"testapi/utils"
)

type ContactService struct {
	store         models.ContactBook
	validationSvc ValidationService
}

func NewContactService(store models.ContactBook, vs ValidationService) *ContactService {

	return &ContactService{
		store:         store,
		validationSvc: vs,
	}
}

func (cs *ContactService) GetAllContacts() ([]models.Contact, error) {

	//DB Call
	cts, err := cs.store.GetContacts()
	return cts, err
}

func (cs *ContactService) CreateContact(c *models.Contact) error {
	if ok := cs.validationSvc.Check(c.Phone); !ok {
		return fmt.Errorf("wrong phone number")
	}

	//DB Call
	err := cs.store.AddContact(c)

	if err != nil {
		return err
	}

	return nil

}

func (cs *ContactService) DeleteContact(c string) error {
	if ok := cs.validationSvc.Check(c); !ok {
		return fmt.Errorf("wrong Phone Number")
	}
	// DB Call
	err := cs.store.DeleteContact(c)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ContactService) UpdateContact(c *models.Contact) error {
	if ok := cs.validationSvc.Check(c.Phone); !ok {
		return fmt.Errorf("wrong phone number")
	}

	// DB Call
	err := cs.store.UpdateContact(c)

	if err != nil {
		return err
	}

	return nil
}

func (cs *ContactService) AddBulkContacts(file io.Reader) error {

	contacts, err := utils.ParseCSV(file)

	if err != nil {
		return err
	}

	errInvalidContact := "Error at: "
	for _, contact := range contacts {
		if ok := cs.validationSvc.Check(contact.Phone); !ok {
			errInvalidContact = errInvalidContact + contact.Name + ", "
		}
	}

	if len(errInvalidContact) > 0 {
		return fmt.Errorf(errInvalidContact)
	}

	//DB Call
	if err = cs.store.InsertManyContacts(contacts); err != nil {
		return err
	}

	return nil
}
