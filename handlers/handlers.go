package handlers

import (
	"encoding/json"
	"net/http"
	"testapi/models"
	"testapi/services"
	"testapi/utils"
)

type ContactHandler struct {
	book models.ContactBook
}

func (h *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	var c models.Contact
	
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Error while decoding", http.StatusInternalServerError)
		return
	}

	if ok := services.NewValidationSevice(c.Phone).Check(); !ok {
		http.Error(w, "Wrong Phone Number", http.StatusBadRequest)
		return
	}


	newErr := h.book.AddContact(&c)


	if newErr != nil {
		http.Error(w, "Error storing data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(utils.EncodeMessage("Contact created"))
}

func (h *ContactHandler) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	cts, _ := h.book.GetContacts()
	jsonBytes, _ := json.Marshal(cts)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	c := q.Get("number")

	newErr := h.book.DeleteContact(c)

	if ok := services.NewValidationSevice(c).Check(); !ok {
		http.Error(w, "Wrong Phone Number", http.StatusBadRequest)
		return
	}

	
	if newErr != nil {
		http.Error(w, "Error deleting Data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(utils.EncodeMessage("Contact Deleted successfully"))
}

func (h *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	var c models.Contact
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Error while decoding", http.StatusInternalServerError)
		return
	}

	if ok := services.NewValidationSevice(c.Phone).Check(); !ok {
		http.Error(w, "Wrong Phone Number", http.StatusBadRequest)
		return
	}


	if updateErr := h.book.UpdateContact(&c); updateErr != nil {
		http.Error(w, "Error updating Data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(utils.EncodeMessage("Contact Updated successfully"))
}


func (h *ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch r.Method {
	case http.MethodPost:
		h.CreateContact(w, r)
		return
	case http.MethodGet:
		h.GetAllContacts(w, r)
		return
	case http.MethodDelete:
		h.DeleteContact(w, r)
		return
	case http.MethodPut:
		h.UpdateContact(w, r)
		return
	}
	
}




// Dependency Injection
func NewContactHandler(book models.ContactBook) *ContactHandler {
	return &ContactHandler{
		book: book,
	}
}



