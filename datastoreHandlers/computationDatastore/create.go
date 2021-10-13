package computationDatastore

import (
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)


func (s *DatastoreStoreWebhook) Create(webhookId int64, values map[string]int64) (int64, error) {

	//Transform map for pub/sub
	var valueToStore []models.CustomMap
	for key, value := range values {
		item := models.CustomMap{
			Key:   key,
			Value: value,
		}
		valueToStore = append(valueToStore, item)
	}

	newKey := datastore.IncompleteKey("Computation", nil)

	computation, err := s.client.Put(context.Background(), newKey,
		&models.ComputationDto{
			WebhookId: webhookId,
			Values:    valueToStore,
			Result:    0,
			Computed:  false,
		})

	if err != nil {
		return 0, err
	}


	return computation.ID, nil

}
