package tests

import (
	mock_datastore "testapi/mocks/datastore"
	"testapi/models"
	"testapi/services"
	"testing"
	"github.com/golang/mock/gomock"
)

func Test_CreateContact(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_datastore.NewMockContactBook(ctrl)
	contact1 := models.Contact{
		Name:  "Test",
		Phone: "9876543210",
	}

	contact2 := models.Contact{
		Name:  "Test",
		Phone: "987654321",
	}

	mockStore.EXPECT().AddContact(&contact1).Return(nil)
	vs := services.NewValidationSevice(10)

	// Act
	cs := services.NewContactService(mockStore, *vs)
	res1 := cs.CreateContact(&contact1)
	res2 := cs.CreateContact(&contact2)

	// Assert
	if res1 != nil {
		t.Error("Error storing data")
	}
	if res2.Error() != "wrong phone number" {
		t.Error(res2)
	}

}

func Test_GetAllContacts(t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_datastore.NewMockContactBook(ctrl)

	mockStore.EXPECT().GetContacts().Return([]models.Contact{models.Contact{Name: "sample", Phone: "99999999"}}, nil)
	vs := services.NewValidationSevice(10)

	// Act
	cs := services.NewContactService(mockStore, *vs)
	res1, _ := cs.GetAllContacts()

	// Assert
	if res1 != nil {
		t.Error("Error Getting data")
	}
}

func Test_DeleteContacts(t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_datastore.NewMockContactBook(ctrl)

	mockStore.EXPECT().DeleteContact("9876543210").Return(nil)
	vs := services.NewValidationSevice(10)

	// Act
	cs := services.NewContactService(mockStore, *vs)

	res := cs.DeleteContact("9876543210")

	// Assert
	if res != nil {
		t.Error(res)
	}
}

func Test_UpdateContact(t *testing.T) {

	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mock_datastore.NewMockContactBook(ctrl)
	contact1 := models.Contact{
		Name:  "Test",
		Phone: "9876543210",
	}
	mockStore.EXPECT().UpdateContact(&contact1).Return(nil)
	vs := services.NewValidationSevice(10)

	// Act
	cs := services.NewContactService(mockStore, *vs)
	res := cs.UpdateContact(&contact1)

	// Assert
	if res != nil {
		t.Error("Error Updating data")
	}
}
