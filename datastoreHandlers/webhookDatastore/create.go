package webhookDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func Create(op string, fields []string) int64 {
	ctx := context.Background()

	client := datastoreHandlers.CreateClient(ctx)

	newKey := datastore.IncompleteKey("Webhook", nil)
	entity, err := client.Put(ctx, newKey,
		&models.WebhookDto{
			Fields: fields,
			Op:     op,
		})

	if err != nil {
		return 0
	}

	defer client.Close()

	return entity.ID

}
