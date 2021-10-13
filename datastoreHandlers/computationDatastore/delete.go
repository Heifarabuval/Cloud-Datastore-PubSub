package computationDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)



func (s *DatastoreStoreWebhook) Delete(id int64) (models.Computation, error) {

	//Verify if computationDatastore exist
	ctx := context.Background()
	dsHandler := InitClient(datastoreHandlers.CreateClient(ctx))
	computation, err := dsHandler.Read(id)
	if err != nil {
		return computation, err
	}


	key := &datastore.Key{
		Kind:      "Computation",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}
	_ = s.client.Delete(ctx, key)

	return computation, nil
}
