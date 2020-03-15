package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

//Insertintouserdb inserts the data into the database
func Insertintouserdb(usercollection *mongo.Collection, u User) bool {

	insertResult, err := usercollection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Print(err)
		return false
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true
}

//Createdb creates a database
func Createdb() (*mongo.Collection, *mongo.Client) {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	usercollection := client.Database("test").Collection("Image")
	return usercollection, client
}

//Update updates the score of the user
func Update(c *mongo.Collection, file string, id int) bool {
	filter := bson.D{
		primitive.E{Key: "id", Value: id},
	}

	update := bson.D{{"$set", bson.D{
		{"img", file},
	},
	},
	}
	updateResult, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Print(err)
		return false
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return true
}

//Findfromuserdb finds the required data
func Findfromuserdb(usercollection *mongo.Collection, id int) bool {
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	var result User

	err := usercollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
