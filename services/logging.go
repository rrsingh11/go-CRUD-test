package services

import (
	// "fmt"

	"io"
	"log"
	"testapi/models"
	"time"
)

type LoggingSvc struct {
	next   models.Service
	logger *log.Logger
}

func (l *LoggingSvc) CreateContact(contact *models.Contact) (err error) {
	defer func(begin time.Time) {
		l.logger.Println("method", "AddContact", "contact", *contact, "error", err, "took", time.Since(begin))
	}(time.Now())
	return l.next.CreateContact(contact)
}
func (l *LoggingSvc) GetAllContacts() (contacts []models.Contact, err error) {

	defer func(begin time.Time) {
		l.logger.Println("method", "GetContacts", "contact", contacts, "error", err, "took", time.Since(begin))
	}(time.Now())

	return l.next.GetAllContacts()
}
func (l *LoggingSvc) DeleteContact(contactName string) (err error) {
	defer func(begin time.Time) {
		l.logger.Println("method", "DeleteContact", "contact", contactName, "error", err, "took", time.Since(begin))
	}(time.Now())
	return l.next.DeleteContact(contactName)
}
func (l *LoggingSvc) UpdateContact(contact *models.Contact) (err error) {
	defer func(begin time.Time) {
		l.logger.Println("method", "UpdateContact", "contact", *contact, "error", err, "took", time.Since(begin))
	}(time.Now())
	return l.next.UpdateContact(contact)
}

func (l *LoggingSvc) AddBulkContacts(r io.Reader) (err error) {
	defer func(begin time.Time) {
		l.logger.Println("method", "AddBulkContacts", "contacts", "" , "error", err, "took", time.Since(begin))
	}(time.Now())
	return l.next.AddBulkContacts(r)
}

func NewLoggingService(serv models.Service, logger *log.Logger) models.Service {
	return &LoggingSvc{
		next:   serv,
		logger: logger,
	}
}
