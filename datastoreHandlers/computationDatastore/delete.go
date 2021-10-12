package computationDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func Delete(id int64) interface{} {

	//Verify if computationDatastore exist
	computation := &models.ComputationRead{}
	deletedComputation := datastoreHandlers.ReadById(id, "Computation", computation)
	if deletedComputation == nil {
		return deletedComputation
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
	return deletedComputation
}
