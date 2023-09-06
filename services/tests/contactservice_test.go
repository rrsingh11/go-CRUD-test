package services

import (
	"testapi/datastore"
	"testapi/models"
	"testapi/services"
	"testing"
)

func Test_CreateContact(t *testing.T) {
	store := datastore.InMemory{Contacts: make(map[string]string)}
	vs := services.NewValidationSevice(10)

	cs := services.NewContactService(&store, *vs)

	err := cs.CreateContact(&models.Contact{
		Name: "Test",
		Phone: "9876543210",
	})

	if err != nil {
		t.Error("Error", err.Error())
	}

	
	if store.Contacts["Test"] != "9876543210" {
		t.Error("Error storing data")
	}

}