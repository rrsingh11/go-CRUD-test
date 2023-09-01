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
		return fmt.Errorf("error while inserting")
	}
	return nil
}

type bsonData []bson.D

func (d bsonData) Decodable() {}

func (m *MongoDB) GetContacts() (models.Output, error) {
	cursor, err := m.Collection.Find(context.Background(), bson.D{})
	// defer cursor.Close(context.Background())

	if err != nil {
		return make(bsonData, 0), fmt.Errorf("error while fetching contacts")
	}

	var contacts bsonData

	if err := cursor.All(context.TODO(), &contacts); err != nil {
		panic(err)
	}
	return contacts, nil
}

func (m *MongoDB) DeleteContact(contactNumber string) error {
	filter := bson.M{"phone": contactNumber}
	_, err := m.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) UpdateContact(contact *models.Contact) error {
	filter := bson.M{"name": contact.Name}
	update := bson.D{{Key: "$set", Value: bson.M{"phone": contact.Phone}}}

	_, err := m.Collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return nil
}
