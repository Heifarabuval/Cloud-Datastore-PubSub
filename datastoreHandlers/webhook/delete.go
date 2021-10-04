package webhook

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func Delete(id int64) interface{} {

	//Verify if webhook exist
	webhook := &models.Webhook{}
	deletedWebhook := datastoreHandlers.ReadById(id, "Webhook", webhook)
	if deletedWebhook == nil {
		return deletedWebhook
	}

	//Create client
	ctx := context.Background()
	client := datastoreHandlers.CreateClient(ctx)


	//Create key for search
	key := &datastore.Key{
		Kind:      "Webhook",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}

	_ = client.Delete(ctx, key)

	defer client.Close()

	return deletedWebhook

}
