package datastore

import (
	"context"
	"fmt"
	"testapi/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	Collection *mongo.Collection
}

func (m *MongoDB) AddContact(contact *models.Contact) error {
	_, err := m.Collection.InsertOne(context.Background(), contact)

	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) GetContacts() ([]models.Contact, error) {
	cursor, err := m.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return make([]models.Contact, 0), fmt.Errorf("error while fetching contacts")
	}
	var contacts []models.Contact

	if err := cursor.All(context.TODO(), &contacts); err != nil {
		return contacts, err
	} else if len(contacts) == 0 {
		return contacts, models.ErrNotFound
	}
	return contacts, nil
}

func (m *MongoDB) DeleteContact(contactNumber string) error {
	filter := bson.M{"phone": contactNumber}
	res, err := m.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return models.ErrNotFound
	}
	return nil
}

func (m *MongoDB) UpdateContact(contact *models.Contact) error {
	filter := bson.M{"name": contact.Name}
	update := bson.D{{Key: "$set", Value: bson.M{"phone": contact.Phone}}}

	res, err := m.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	} else if res.MatchedCount == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (m *MongoDB) InsertManyContacts(contacts []models.Contact) error {
	var cts []interface{}
	for _, contact := range contacts {
		cts = append(cts, contact)
	}

	_, err := m.Collection.InsertMany(context.Background(), cts)

	if err != nil {
		return err
	}
	return nil
}
