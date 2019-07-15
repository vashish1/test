package datastore

type Datastore interface {
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
	Close()
}



func NewDatastore(dbConnectionString string) (Datastore, error) {

	
				return NewMongoDBDatastore(dbConnectionString)
			
	}

	return nil, nil
}
