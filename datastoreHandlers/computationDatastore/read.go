package computationDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
)



func Read(id int64) interface{} {
	computation := &models.ComputationRead{}
	exist := datastoreHandlers.ReadById(id, "Computation", computation)

	if exist == nil {
		return exist
	}


	computationFinal := models.Computation{
		ID:        computation.ID,
		WebhookId: computation.WebhookId,
		Result:    computation.Result,
		Values:    computation.TransformToMap(computation.Values),
		Computed:  computation.Computed,
	}

	return computationFinal
}
