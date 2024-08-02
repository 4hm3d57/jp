package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Profile struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	First_name   string             `bson:"f_name"`
	Last_name    string             `bson:"l_name"`
	Born         string             `bson:"born"`
	Email        string             `bson:"email"`
	Ed_Level     string             `bson:"ed_level"`
	Ed_course    string             `bson:"ed_course"`
	Gender       string             `bson:"gender"`
	City         string             `bson:"city"`
	Street       string             `bson:"street"`
	Zip          string             `bson:"zip"`
	County       string             `bson:"county"`
	Phone_number string             `bson:"phone"`
	About        string             `bson:"about"`
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

func GetProfile(name, email string) (*User, error) {

	client, userCollection, err := UserDB()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	var user User
	err = userCollection.FindOne(context.Background(), bson.M{"name": name, "email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
