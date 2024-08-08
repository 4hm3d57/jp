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
	Name        string             `bson:"name"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	AccountType string             `bson:"acc_type"`
}

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

type Contacts struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name"`
	Email   string             `bson:"email"`
	Message string             `bson:"message"`
}

type Academics struct {
	ID          primitive.ObjectID
	Ed_level    string
	Institution string
	Course      string
	Timeframe   string
}

type Experience struct {
	ID          primitive.ObjectID
	Institution string
	Supervisor  string
	Telephone   string
	Jobtitle    string
	Start       string
	End         string
	Duties      string
}

type Language struct {
	ID    primitive.ObjectID
	Lang  string
	Speak string
	Read  string
	Write string
}

type Profession struct {
	ID          primitive.ObjectID
	County      string
	Institution string
	Course      string
	Timeframe   string
}

type Referee struct {
	ID          primitive.ObjectID
	Name        string
	Email       string
	Title       string
	Phone       string
	Institution string
}

type Train struct {
	ID          primitive.ObjectID
	Training    string
	Institution string
	Timeframe   string
}

type EmployerProfile struct {
	ID          primitive.ObjectID
	Name        string
	Established string
	Type        string
	People      string
	Website     string
	City        string
	Street      string
	Zip         string
	Phone       string
	Email       string
	Background  string
	Service     string
	Expertise   string
}

// db connection functions
func UserDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	userCollection := client.Database("jp").Collection("users")

	return client, userCollection, nil
}

func ProfileDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	profileCollection := client.Database("jp").Collection("profiles")

	return client, profileCollection, nil

}

func ContactsDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	userCollection := client.Database("jp").Collection("contacts")

	return client, userCollection, nil
}

func AcademicsDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	academicsCollection := client.Database("jp").Collection("academics")

	return client, academicsCollection, nil
}

func ExpreinceDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	expCollection := client.Database("jp").Collection("experience")

	return client, expCollection, nil

}

func LanguageDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	langCollection := client.Database("jp").Collection("language")

	return client, langCollection, nil
}

func ProfessionDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	profCollection := client.Database("jp").Collection("profession")

	return client, profCollection, nil

}

func RefereeDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	refCollection := client.Database("jp").Collection("referees")

	return client, refCollection, nil

}

func TrainDB() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	trainCollection := client.Database("jp").Collection("training")

	return client, trainCollection, nil

}

func EmpProfile() (*mongo.Client, *mongo.Collection, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, nil, err
	}

	empCollection := client.Database("jp").Collection("employer_profile")

	return client, empCollection, nil

}

// users db

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

func GetUserID(userID primitive.ObjectID) (*User, error) {

	client, userCollection, err := UserDB()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	var user User
	err = userCollection.FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func GetUserLogin(email, password string) (*User, error) {

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

func GetUserSignup(name, email string) (*User, error) {

	client, userCollection, err := UserDB()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	var profile User
	err = userCollection.FindOne(context.Background(), bson.M{"name": name, "email": email}).Decode(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func GetAllUser() ([]User, error) {

	client, userCollection, err := UserDB()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	var user []User
	cur, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(context.Background(), &user); err != nil {
		return nil, err
	}

	return user, nil

}

func GetPass(password string) (*User, error) {

	client, userCollection, err := UserDB()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	var pass User
	err = userCollection.FindOne(context.Background(), bson.M{"password": password}).Decode(&pass)
	if err != nil {
		return nil, err
	}

	return &pass, nil

}

func UpdatePass(id primitive.ObjectID, newPassword string) error {

	client, userCollection, err := UserDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	_, err = userCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"password": newPassword}},
	)

	if err != nil {
		return err
	}

	return nil
}

// profile db

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

func GetProfileID(id primitive.ObjectID) (*Profile, error) {
	client, userCollection, err := ProfileDB()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	var profile Profile
	filter := bson.M{"_id": id}
	err = userCollection.FindOne(context.Background(), filter).Decode(&profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func UpdateProfile(profile Profile) error {

	client, userCollection, err := ProfileDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	filter := bson.M{"_id": profile.ID}
	update := bson.M{
		"$set": bson.M{
			"first_name":   profile.First_name,
			"last_name":    profile.Last_name,
			"born":         profile.Born,
			"email":        profile.Email,
			"ed_level":     profile.Ed_Level,
			"ed_course":    profile.Ed_course,
			"gender":       profile.Gender,
			"city":         profile.City,
			"street":       profile.Street,
			"zip":          profile.Zip,
			"county":       profile.County,
			"phone_number": profile.Phone_number,
			"about":        profile.About,
		},
	}
	_, err = userCollection.UpdateOne(context.Background(), filter, update)
	return err
}

// contact db

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

// academics db

func InsertAcademics(academics Academics) error {

	client, academicsCollection, err := AcademicsDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	_, err = academicsCollection.InsertOne(context.Background(), academics)
	if err != nil {
		return err
	}

	return nil
}

// experience db

func InsertExperience(experience Experience) error {

	client, expCollection, err := ExpreinceDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	_, err = expCollection.InsertOne(context.Background(), experience)
	if err != nil {
		return err
	}

	return nil
}

// language db

func InsertLanguage(language Language) error {

	client, langCollection, err := LanguageDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	_, err = langCollection.InsertOne(context.Background(), language)
	if err != nil {
		return err
	}

	return nil
}

// profession db

func InsertProfession(profession Profession) error {

	client, profCollection, err := ProfessionDB()
	if err != nil {
		return nil
	}
	defer client.Disconnect(context.Background())

	_, err = profCollection.InsertOne(context.Background(), profession)
	if err != nil {
		return err
	}

	return nil
}

// referee db

func InsertReferee(referee Referee) error {

	client, refCollection, err := RefereeDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	_, err = refCollection.InsertOne(context.Background(), referee)
	if err != nil {
		return err
	}

	return nil
}

// training db

func InsertTrainData(train Train) error {

	client, trainCollection, err := TrainDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	_, err = trainCollection.InsertOne(context.Background(), train)
	if err != nil {
		return err
	}

	return nil
}

// employee profile

func InsertEmployerProfile(profile EmployerProfile) error {

	client, empCollection, err := EmpProfile()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	_, err = empCollection.InsertOne(context.Background(), profile)
	if err != nil {
		return err
	}

	return nil
}
