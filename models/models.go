package models

import (
	"encoding/json"
	"io"
)

// ContactInfo
type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// Service Port
type Service interface {
	DeleteContact(string) error
	UpdateContact(*Contact) error
	CreateContact(*Contact) error
	GetAllContacts() ([]Contact, error)
	AddBulkContacts(io.Reader) error
}

// Database port


type Config struct {
	Memory struct {
		Type       string `json:"type"`
		Database   string `json:"database"`
		Collection string `json:"collection"`
	} `json:"memory"`
	Host string `json:"host"`
	Port int    `json:"port"`
	Test_Memory struct { 
		Type       string `json:"type"`
		Database   string `json:"database"`
		Collection string `json:"collection"`
	} `json:"test-memory"` 
}

type Reponse struct {
	Status  string          `json:"status,omitempty"`
	Message string          `json:"message,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
	Code    string          `json:"code,omitempty"`
}

type BError struct {
	Message string
	Code    int
}

func (b BError) Error() string {
	return b.Message
}
