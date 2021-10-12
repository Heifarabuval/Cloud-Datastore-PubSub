package webhookDatastore

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
)

func Read(id int64) *models.Webhook {

	webhook := &models.Webhook{}
	exist := datastoreHandlers.ReadById(id, "Webhook", webhook)

	if exist == nil {
		return webhook
	}

	return webhook

}
