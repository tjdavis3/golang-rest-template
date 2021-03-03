package models

// Datastore is an interface to the backend datastore
type Datastore interface {
}

// InitializeDB creates a DB connection from the provided configuration
func Initialize(config interface{}) (Datastore, error) {
	return nil, nil
}
