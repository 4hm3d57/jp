package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Profile struct {
	ID           primitive.ObjectID
	First_name   string
	Last_name    string
	Born         string
	Email        string
	Ed_Level     string
	Ed_course    string
	Gender       string
	City         string
	Street       string
	Zip          string
	County       string
	Phone_number string
	About        string
}

func ProfileDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	profileCollection := client.Database("jp").Collection("profiles")

	return client, profileCollection, nil

}

func InsertProfile(profile Profile) error {

	client, profileCollection, err := ProfileDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	profile.ID = primitive.NewObjectID()
	_, err = profileCollection.InsertOne(context.Background(), profile)
	if err != nil {
		return err
	}

	return nil

}
