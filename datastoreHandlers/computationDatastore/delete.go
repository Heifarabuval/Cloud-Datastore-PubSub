package computationDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func Delete(id int64) (models.ComputationRead, error) {

	//Verify if computationDatastore exist
	deletedComputation, err := Read(id)
	if err != nil {
		return deletedComputation, err
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
	return deletedComputation, nil
}
