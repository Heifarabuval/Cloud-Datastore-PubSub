package datastoreHandlers

import (
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
	"log"
)
func Create( op string, fields []string) int64 {
	ctx := context.Background()

	projectID := "heifara-test"

	client, err := datastore.NewClient(ctx,projectID)

	if err != nil{
		log.Fatalf("Failed to create client: %v",err)
	}


	newKey := datastore.IncompleteKey("Webhook", nil)
	test, error := client.Put(ctx, newKey,
		&models.WebhookDto{
			Fields: fields,
			Op:     op,
		})


	if error != nil{
		log.Fatalf("Failed to create client: %v",err)
	}


	defer client.Close()

	return test.ID


}

