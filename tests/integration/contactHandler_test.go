package tests

import (
	"fmt"
	"testapi/datastore"
	"testapi/models"
	"testapi/services"
	"testapi/utils"
	"testing"
)

func testSetup() models.Service {
	// Setup
	config, err := utils.LoadConfiguration("../../config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	mongoCollection := utils.ConnectMongoDB(config.Test_Memory.Database, config.Test_Memory.Collection)
	store := &datastore.MongoDB{Collection: mongoCollection}

	vs := services.NewValidationSevice(10)
	//Contact service
	cs := services.NewContactService(store, *vs)

	return cs
}

func Test_CreateContact_ContactService_MongoDB(t *testing.T) {
	cs := testSetup()
	var test_contact1 = models.Contact{Name: "Test1", Phone: "1234567890"}

	// Create Contact
	err := cs.CreateContact(&test_contact1)
	if err != nil {
		t.Error(err.Error())
	}
	cts, err := cs.GetAllContacts()
	if err != nil || len(cts) == 0 || cts[0] != test_contact1 {
		t.Error("Create Contact not working")
	}

}

func Test_UpdateContact_ContactService_MongoDB(t *testing.T) {
	cs := testSetup()
	var test_contact2 = models.Contact{Name: "Test1", Phone: "1234567891"}

	// Update Contact
	err := cs.UpdateContact(&test_contact2)

	if err != nil {
		t.Error(err.Error())
	}
	if err == nil {
		cts, _ := cs.GetAllContacts()
		if cts[0] != test_contact2 {
			t.Error("Update Contact not working")
		}
	}
}

func Test_DeleteContact_ContactService_MongoDB(t *testing.T) {
	cs := testSetup()
	// Delete Contact
	err := cs.DeleteContact("1234567891")
	if err != nil {
		t.Error(err.Error())
	}

	_, err = cs.GetAllContacts()

	if err.Error() != "Records Not Found" {
		t.Error("Delete Contact not working", err.Error())
	}
}
