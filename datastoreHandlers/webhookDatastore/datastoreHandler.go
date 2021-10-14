package webhookDatastore

import "cloud.google.com/go/datastore"

type DatastoreStoreWebhook struct {
	client *datastore.Client
}

func InitClient(client *datastore.Client) DatastoreStoreWebhook {
	return DatastoreStoreWebhook{client: client}
}
