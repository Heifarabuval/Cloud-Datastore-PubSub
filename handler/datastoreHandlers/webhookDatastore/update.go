package webhookDatastore

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/handler/datastoreHandlers"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
)

func (s *DatastoreStoreWebhook) Update(id int64, op string, fields []string) (models.Webhook, error) {
	ctx := context.Background()
	dsHandler := InitClient(datastoreHandlers.CreateClient(ctx))
	webhook, err := dsHandler.Read(id)

	if err != nil {
		return webhook, nil
	}
	if len(op) == 0 {
		op = webhook.Op
	}
	if len(fields) == 0 {
		fields = webhook.Fields
	}

	key := &datastore.Key{
		Kind:      "Webhook",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}

	_, err = s.client.Put(ctx, key,
		&models.WebhookDto{
			Fields: fields,
			Op:     op,
		})

	webhook, err = s.Read(id)
	if err != nil {
		return webhook, nil
	}

	if err != nil {
		return webhook, nil
	}

	return webhook, nil
}
