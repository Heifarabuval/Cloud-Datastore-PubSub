package webhookDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func Read(id int64) (models.Webhook, error) {

	webhook := models.Webhook{}

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

	err := client.Get(ctx, key, &webhook)

	if err != nil {
		return webhook, err
	}

	defer client.Close()

	return webhook, nil

}
