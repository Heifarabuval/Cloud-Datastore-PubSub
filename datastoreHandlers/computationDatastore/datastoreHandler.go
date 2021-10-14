package computationDatastore

import (
	"cloud.google.com/go/datastore"
)

type DatastoreStoreComputation struct {
	client *datastore.Client
}

func InitClient(client *datastore.Client) DatastoreStoreComputation {
	return DatastoreStoreComputation{client: client}
}
