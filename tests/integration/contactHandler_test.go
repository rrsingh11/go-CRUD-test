package tests

import (
	"testapi/datastore"
	"testapi/models"
	"testapi/services"
	"testapi/utils"
	"testing"
)

func Test_ContactService_MongoDB(t *testing.T) {
	// Setup
	config, _ := utils.LoadConfiguration("config.json")
	mongoCollection := utils.ConnectMongoDB(config.Test_Memory.Database, config.Test_Memory.Collection)
	store := &datastore.MongoDB{Collection: mongoCollection}

	vs := services.NewValidationSevice(10)
	//Contact service
	cs := services.NewContactService(store, *vs)

	var test_contact = models.Contact{Name: "Test", Phone: "1234567890"}
	err := cs.CreateContact(&test_contact)
	if err == nil {
		cts, _ := cs.GetAllContacts()
		// fmt.Println(cts)
		if cts[0] != test_contact {
			t.Error("Create Contact not working")

		}
	}

}
