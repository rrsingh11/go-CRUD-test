package testapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"testapi/datastore"
	"testapi/models"
)


type ContactHandler struct {
	*sync.Mutex
	book datastore.ContactBook
}

func (h *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	var c models.Contact
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	if err != nil {
		fmt.Println("Error while decoding", err)
		return
	}

	newErr := h.book.AddContact(&c)
	if newErr != nil {
		fmt.Fprintln(w, "Trouble storing data")
		return
	}
	fmt.Fprintf(w, "Contact added Successfully")
}

func (h *ContactHandler) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	cts, _ := h.book.GetContacts()
	jsonBytes, _ := json.Marshal(cts)
	w.Write(jsonBytes)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	c := q.Get("number")
	newErr := h.book.DeleteContact(c)
	if newErr != nil {
		fmt.Fprintln(w, "Trouble deleting data")
		return
	}
	fmt.Fprintln(w, "Contact deleted Successfully")
}

func (h *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	var c models.Contact
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	if err != nil {
		fmt.Println("Error while decoding", err)
		return
	}

	newErr := h.book.UpdateContact(&c)
	if newErr != nil {
		fmt.Fprintln(w, "Trouble updating data")
		return
	}
	fmt.Fprintf(w, "Contact Updated Successfully")
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





func NewHandler() ContactHandler {
	return ContactHandler{
		book: &datastore.InMemory{
			Contacts: make(map[string]string),
			Mutex: &sync.Mutex{},
		},
	}
}
