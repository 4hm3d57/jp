package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string `bson:"name"`
	Email       string `bson:"email"`
	Password    string `bson:"password"`
	AccountType string `bson:"acc_type"`
}

func UserDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	userCollection := client.Database("jp").Collection("users")

	return client, userCollection, nil
}

func InsertUser(user User) error {

	client, userCollection, err := UserDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	user.ID = primitive.NewObjectID()
	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(email, password string) (*User, error) {

	client, userCollection, err := UserDB()
	if err != nil {
		log.Print("Failed to connect to db")
		return nil, err
	}
	defer client.Disconnect(context.Background())

	var user User
	err = userCollection.FindOne(context.Background(), bson.M{"email": email, "password": password}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
