package computationDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func Delete(id int64) (models.ComputationRead, error) {

	//Verify if computationDatastore exist
	computation, err := Read(id)
	if err != nil {
		return computation, err
	}

	ctx := context.Background()
	client := datastoreHandlers.CreateClient(ctx)

	key := &datastore.Key{
		Kind:      "Computation",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}
	_ = client.Delete(ctx, key)
	defer client.Close()
	return computation, nil
}
