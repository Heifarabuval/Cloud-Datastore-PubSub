package webhookDatastore

import (
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/datastoreHandlers"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
	"cloud.google.com/go/datastore"
	"context"
)

func (s *DatastoreStoreWebhook) Delete(id int64) (models.Webhook, error) {

	//Verify if webhookDatastore exist
	dsHandler := InitClient(datastoreHandlers.CreateClient(context.Background()))
	webhook, err := dsHandler.Read(id)
	if err != nil {
		return webhook, err
	}

	//Create key for search
	key := &datastore.Key{
		Kind:      "Webhook",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}

	_ = s.client.Delete(context.Background(), key)

	return webhook, nil

}
