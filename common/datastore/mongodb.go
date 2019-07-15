package datastore

import (
	"fmt"

	"test/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MongoDBDatastore .....
type MongoDBDatastore struct {
	*mgo.Session
}

//NewMongoDBDatastore ......
func NewMongoDBDatastore(url string) (*MongoDBDatastore, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &MongoDBDatastore{
		Session: session,
	}, nil
}

//CreateUser .....
func (m *MongoDBDatastore) CreateUser(user *models.User) error {

	session := m.Copy()

	defer session.Close()
	userCollection := session.DB("gopherface").C("User")
	err := userCollection.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

//GetUser ....
func (m *MongoDBDatastore) GetUser(username string) (*models.User, error) {

	session := m.Copy()
	defer session.Close()
	userCollection := session.DB("gopherface").C("User")
	u := models.User{}
	err := userCollection.Find(bson.M{"username": username}).One(&u)
	if err != nil {
		return nil, err
	}
	fmt.Println("created")
	return &u, nil

}

//Close ........
func (m *MongoDBDatastore) Close() {
	m.Close()
}
