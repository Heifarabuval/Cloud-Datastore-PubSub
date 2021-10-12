package computationDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func ReadAll() []models.Computation {
	ctx := context.Background()
	client := datastoreHandlers.CreateClient(ctx)

	//Model
	var computations []models.ComputationRead
	var computationsFinal []models.Computation

	//Request
	_, err := client.GetAll(ctx, datastore.NewQuery("Computation"), &computations)

	for i, v := range computations {

		computation := models.Computation{
			ID:        v.ID,
			WebhookId: v.WebhookId,
			Result:    v.Result,
			Values:    computations[i].TransformToMap(computations[i].Values),
			Computed:  v.Computed,
		}
		computationsFinal = append(computationsFinal, computation)
	}

	if err != nil {
		return nil
	}

	defer client.Close()

	return computationsFinal

}
