package webhookDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func Update(id int64, op string, fields []string) (models.Webhook, error) {
	ctx := context.Background()
	client := datastoreHandlers.CreateClient(ctx)

	webhook, err := Read(id)
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

	_, err = client.Put(ctx, key,
		&models.WebhookDto{
			Fields: fields,
			Op:     op,
		})

	webhook, err = Read(id)
	if err != nil {
		return webhook, nil
	}

	if err != nil {
		return webhook, nil
	}

	return webhook, nil
}
