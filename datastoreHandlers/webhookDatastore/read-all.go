package webhookDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func ReadAll() []models.Webhook {
	ctx := context.Background()
	client := datastoreHandlers.CreateClient(ctx)

	//Model
	var webhooks []models.Webhook

	//Request
	_, err := client.GetAll(ctx, datastore.NewQuery("Webhook"), &webhooks)

	if err != nil {
		return nil
	}

	defer client.Close()

	return webhooks

}
