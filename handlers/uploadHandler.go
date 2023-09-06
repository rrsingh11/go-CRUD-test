package handlers

import (
	"bytes"
	"encoding/csv"
	"io"

	// "io"
	"net/http"
	// "strconv"
	"testapi/models"
	"testapi/services"
	"testapi/utils"
)

type UploadHandler struct {
	Services services.Services
}

func (up UploadHandler) UploadFile(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20) //max 10MB
	file, _, err := r.FormFile("filename")

	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if err = up.Services.Contactservice.AddBulkContacts(file); err != nil {
		utils.EncodeError(w, err)
		return
	}

	w.Write(utils.EncodeMessage(models.Reponse{
		Status:  "success",
		Code:    "S",
		Message: "Contacts Uploaded Successfully",
	}))
}

func (up UploadHandler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	cts, err := up.Services.Contactservice.GetAllContacts()
	if err != nil {
		utils.EncodeError(w, err)
		return
	}

	buff := new(bytes.Buffer)
	//csv writer
	csvWriter := csv.NewWriter(buff)
	header := []string{"Name", "Phone"}
	err = csvWriter.Write(header)

	if err != nil {
		utils.EncodeError(w, err)
		return
	}

	//Data to csv
	for _, contact := range cts {
		record := []string{contact.Name, contact.Phone}
		if err = csvWriter.Write(record); err != nil {
			utils.EncodeError(w, err)
		}

	}
	csvWriter.Flush()

	if err = csvWriter.Error(); err != nil {
		utils.EncodeError(w, err)
		return
	}

	io.Copy(w, buff)
}

func (up UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	switch r.Method {
	case http.MethodPost:
		up.UploadFile(w, r)
	case http.MethodGet:
		up.DownloadFile(w, r)
	}

}

func NewUploadHandler(serv services.Services) UploadHandler {
	return UploadHandler{
		Services: serv,
	}
}
