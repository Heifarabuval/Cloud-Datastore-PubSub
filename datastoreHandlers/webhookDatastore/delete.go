package webhookDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func Delete(id int64) (models.Webhook, error) {

	//Verify if webhookDatastore exist
	webhook, err := Read(id)
	if err != nil {
		return webhook, err
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

	return webhook, nil

}
