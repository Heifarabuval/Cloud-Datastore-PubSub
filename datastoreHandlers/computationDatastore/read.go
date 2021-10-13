package computationDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func Read(id int64) (models.ComputationRead, error) {

	computation := models.ComputationRead{}
	ctx := context.Background()
	client := datastoreHandlers.CreateClient(ctx)

	//Create key for search
	key := &datastore.Key{
		Kind:      "Computation",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}

	err := client.Get(ctx, key, &computation)

	if err != nil {
		return computation, err
	}

	defer client.Close()

	computationPs := models.Computation{}

	//Hydrate data
	computationPs.ID = computation.ID
	computationPs.WebhookId = computation.WebhookId
	computationPs.Result = computation.Result
	computationPs.Values = computation.TransformToMap(computation.Values)
	computationPs.Computed = computation.Computed
	return computation, nil
}
