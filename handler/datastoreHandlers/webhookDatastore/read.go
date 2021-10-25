package webhookDatastore

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
)

func (s *DatastoreStoreWebhook) Read(id int64) (models.Webhook, error) {

	webhook := models.Webhook{}

	//Create key for search
	key := &datastore.Key{
		Kind:      "Webhook",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}

	err := s.client.Get(context.Background(), key, &webhook)

	if err != nil {
		return webhook, err
	}

	return webhook, nil

}
