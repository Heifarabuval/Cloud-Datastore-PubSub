package webhook

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
)

func Read(id int64) interface{} {

	webhook := &models.Webhook{}
	exist := datastoreHandlers.ReadById(id, "Webhook", webhook)
	if exist == nil {
		return exist
	}
	return webhook

}
