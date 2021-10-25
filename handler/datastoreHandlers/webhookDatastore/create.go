package webhookDatastore

import (
	"context"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"

	"cloud.google.com/go/datastore"
)

func (s *DatastoreStoreWebhook) Create(op string, fields []string) (int64, error) {

	newKey := datastore.IncompleteKey("Webhook", nil)
	webhook, err := s.client.Put(context.Background(), newKey,
		&models.WebhookDto{
			Fields: fields,
			Op:     op,
		})

	if err != nil {
		return 0, err
	}

	return webhook.ID, nil

}
