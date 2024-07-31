package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Contacts struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Email   string             `bson:"email"`
	Message string             `bson:"message"`
}

func ContactsDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	userCollection := client.Database("jp").Collection("contacts")

	return client, userCollection, nil
}

func InsertContact(contacts Contacts) error {

	client, userCollection, err := ContactsDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	contacts.ID = primitive.NewObjectID()
	_, err = userCollection.InsertOne(context.Background(), contacts)
	if err != nil {
		return err
	}

	return nil
}
