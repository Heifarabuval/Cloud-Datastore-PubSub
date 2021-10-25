package computationDatastore

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
)

func (s *DatastoreStoreComputation) ReadAll() ([]models.Computation, error) {

	//Model
	var computations []models.ComputationRead
	var computationsFinal []models.Computation

	//Request
	_, err := s.client.GetAll(context.Background(), datastore.NewQuery("Computation"), &computations)

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
		return []models.Computation{}, err
	}

	return computationsFinal, nil

}
