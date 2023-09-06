package utils

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"testapi/models"
)

func EncodeMessage(resp models.Reponse) []byte {
	res, _ := json.Marshal(resp)
	return res
}

func EncodeError(w http.ResponseWriter, err error) {
	var berr models.BError
	if errors.As(err, &berr) {
		w.WriteHeader(http.StatusNotFound)
		w.Write(EncodeMessage(models.Reponse{
			Status:  "failure",
			Code:    "E",
			Message: err.Error(),
		}))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(EncodeMessage(models.Reponse{
			Status:  "failure",
			Code:    "E",
			Message: err.Error(),
		}))
	}
}

func LoadConfiguration(file string) (models.Config, error) {
	var config models.Config
	configFile, err := os.Open(file)
	if err != nil {
		return *new(models.Config), err
	}

	json.NewDecoder(configFile).Decode(&config)
	return config, nil
}

func ParseCSV(file io.Reader) ([]models.Contact, error) {
	var contacts []models.Contact
	csvReader := csv.NewReader(file)

	data, readErr := csvReader.ReadAll()

	if readErr != nil {
		return nil, readErr
	}

	for i, line := range data {
		if i == 0 {
			continue
		}
		contacts = append(contacts, models.Contact{Name: line[0], Phone: line[1]})
	}

	return contacts, nil
}
