package datastore

import "test/models"

//Datastore ....
type Datastore interface {
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
	Close()
}

//NewDatastore ....
func NewDatastore(dbConnectionString string) (Datastore, error) {

	return NewMongoDBDatastore(dbConnectionString)

}
