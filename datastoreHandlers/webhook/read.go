package datastoreHandlers

import (
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
	"log"
)
func Read(id int64) *models.Webhook {
	ctx := context.Background()

	projectID := "heifara-test"


	client, err := datastore.NewClient(ctx,projectID)

	if err != nil{
		log.Fatalf("Failed to create client: %v",err)
	}

	key := &datastore.Key{
		Kind:      "Webhook",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}


	 webhook := new (models.Webhook)

	_ = client.Get(ctx, key, webhook)

	defer client.Close()

	return webhook


}

