package computationDatastore

import (
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)



func (s *DatastoreStoreWebhook)  Read(id int64) (models.Computation, error) {

	computationDs := models.ComputationRead{}

	//Create key for search
	key := &datastore.Key{
		Kind:      "Computation",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}

	err := s.client.Get(context.Background(), key, &computationDs)
	computation := models.Computation{}

	if err != nil {
		return computation, err
	}


	//Hydrate data
	computation.ID = computationDs.ID
	computation.WebhookId = computationDs.WebhookId
	computation.Result = computationDs.Result
	computation.Values = computationDs.TransformToMap(computationDs.Values)
	computation.Computed = computationDs.Computed
	return computation, nil
}
