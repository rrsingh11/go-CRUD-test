package handlers

import (
	"encoding/json"
	"net/http"
	"testapi/models"
	"testapi/services"
	"testapi/utils"
)

type ContactHandler struct {
	Services services.Services
}

func (h *ContactHandler) PostRequest(w http.ResponseWriter, r *http.Request) {
	var c models.Contact
	// req
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Error while decoding", http.StatusInternalServerError)
		return
	}

	// service call
	newErr := h.Services.Contactservice.CreateContact(&c)

	// response
	if newErr != nil {
		utils.EncodeError(w, newErr)
		return
	}

	w.Write(utils.EncodeMessage(models.Reponse{
		Message: "Deleted Successfully", 
		Code: "S",
		Status: "success",
	}))
}

func (h *ContactHandler) GetRequest(w http.ResponseWriter, r *http.Request) {
	//service call
	cts, err :=	h.Services.Contactservice.GetAllContacts()

	// response
	if err != nil {
		utils.EncodeError(w, err)
		return
	}

	jsonBytes, _ := json.Marshal(cts)

	w.Write(utils.EncodeMessage(models.Reponse{
		Message: "Data recieved Successfully", 
		Code: "S",
		Status: "success",
		Data: jsonBytes,
	}))
}

func (h *ContactHandler) DeleteRequest(w http.ResponseWriter, r *http.Request) {
	// request
	q := r.URL.Query()
	c := q.Get("number")

	// service call
	err := h.Services.Contactservice.DeleteContact(c)

	// response
	if err != nil {
		utils.EncodeError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(utils.EncodeMessage(models.Reponse{
		Message: "Deleted Successfully", 
		Code: "S",
		Status: "success",
	}))
}

func (h *ContactHandler) PutRequest(w http.ResponseWriter, r *http.Request) {
	var c models.Contact
	// request
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Error while decoding", http.StatusInternalServerError)
		return
	}

	// service call
	err := h.Services.Contactservice.UpdateContact(&c)
	
	// resposne
	if err != nil {
		utils.EncodeError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(utils.EncodeMessage(models.Reponse{
		Message: "Updated Successfully", 
		Code: "S",
		Status: "success",
	}))
}


func (h *ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch r.Method {
	case http.MethodPost:
		h.PostRequest(w, r)
		return
	case http.MethodGet:
		h.GetRequest(w, r)
		return
	case http.MethodDelete:
		h.DeleteRequest(w, r)
		return
	case http.MethodPut:
		h.PutRequest(w, r)
		return
	}
	
}




// Dependency Injection
func NewContactHandler(services services.Services) *ContactHandler {
	return &ContactHandler{
		Services: services,
	}
}



