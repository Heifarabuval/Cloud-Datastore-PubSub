package computationDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func Create(webhookId int64, values map[string]int64) int64 {
	ctx := context.Background()
	var valueToStore []models.CustomMap
	for key, value := range values {
		item := models.CustomMap{
			Key:   key,
			Value: value,
		}
		valueToStore = append(valueToStore, item)
	}

	client := datastoreHandlers.CreateClient(ctx)

	newKey := datastore.IncompleteKey("Computation", nil)
	entity, err := client.Put(ctx, newKey,
		&models.ComputationDto{
			WebhookId: webhookId,
			Values:    valueToStore,
			Result:    0,
			Computed:  false,
		})

	if err != nil {
		return 0
	}

	defer client.Close()

	return entity.ID

}
