package webhookDatastore

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
)

func (s *DatastoreStoreWebhook) ReadAll() []models.Webhook {

	//Model
	var webhooks []models.Webhook

	//Request
	_, err := s.client.GetAll(context.Background(), datastore.NewQuery("Webhook"), &webhooks)

	if err != nil {
		return nil
	}

	return webhooks

}
