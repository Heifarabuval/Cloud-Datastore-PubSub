package datastoreHandlers

import (
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"log"
)
func ReadAll() []models.Webhook {
	ctx := context.Background()

	projectID := "heifara-test"

	client, err := datastore.NewClient(ctx,projectID)

	if err != nil{
		log.Fatalf("Failed to create client: %v",err)
	}
	var webhooks []models.Webhook

	keys, err := client.GetAll(ctx, datastore.NewQuery("Webhook"), &webhooks)
	fmt.Println(&webhooks)
	if err != nil {
		return nil
	}
	fmt.Println(keys)


	defer client.Close()

	return webhooks

}

