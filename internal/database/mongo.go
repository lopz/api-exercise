package database

import (
	"context"
	"errors"
	_ "fmt"
	"log"
	"os"
	"time"

	"github.com/lopz/cs-api-test/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var db mongo.Client

var coll *mongo.Collection

var uri = os.Getenv("MONGODB_URI")

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		println("Can't connect to MongoDB!")
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	println("Connect to MongoDB!")

	coll = client.Database("csapi").Collection("people")

	/* 	defer func() {
	   		println("Disconnect from MongoDB")

	   		if err = client.Disconnect(ctx); err != nil {
	   			panic(err)
	   		}
	   	}()
	*/
}

func DbGetAllPerson() ([]*models.Person, error) {
	//person = &models.Person{}
	cur, err := coll.Find(context.TODO(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var people []*models.Person

	for cur.Next(context.TODO()) {
		person := &models.Person{}
		er := cur.Decode(person)
		if er != nil {
			panic(er)
		}
		people = append(people, person)
	}

	return people, errors.New("person not found")
}

func DbAddPerson(person models.Person) (string, error) {

	_, err := coll.InsertOne(context.TODO(), person)

	if err != nil {
		log.Fatal(err)
	}
	return "res.InsertedID", errors.New("person not add")
}

func DbGetPerson(uuid string) (models.Person, error) {
	var person models.Person
	err := coll.FindOne(context.TODO(), bson.M{"uuid": uuid}).Decode(&person)

	if err == mongo.ErrNoDocuments {
		println("record does not exist")
	} else if err != nil {
		panic(err)
	}

	return person, errors.New("person not found")
}

func DbUpdatePerson(uuid string, person models.Person) (models.Person, error) {

	person.UUID = uuid

	filter := bson.M{"uuid": uuid}
	update := bson.M{"$set": person}
	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if result.ModifiedCount != 0 {
		println("No document updated")
	} else if err != nil {
		panic(err)
	}
	person, _ = DbGetPerson(uuid)
	return person, errors.New("person not found")
}

func DbDeletePerson(uuid string) (models.Person, error) {
	person, _ := DbGetPerson(uuid)

	filter := bson.M{"uuid": uuid}
	result, err := coll.DeleteOne(context.TODO(), filter)

	if result.DeletedCount == 0 {
		println("No document deleted")
	} else if err != nil {
		panic(err)
	}
	return person, errors.New("person not found")
}
