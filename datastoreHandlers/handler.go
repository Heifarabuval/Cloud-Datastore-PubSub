package datastoreHandlers

import (
	"Calicut/models"
	"cloud.google.com/go/datastore"
)


type DatastoreStoreWebhook struct {
	client *datastore.Client
}

type StoreWebhook interface {
	Create(op string, fields []string) int64
	Read(id int64) (models.Webhook, error)
	ReadAll() []models.Webhook
	Update(id int64, op string, fields []string) (models.Webhook, error)
	Delete(id int64) (models.Webhook, error)
}

type StoreComputation interface {
	Create(webhookId int64, values map[string]int64) (int64, error)
	Read(id int64) (models.ComputationRead, error)
	ReadAll() ([]models.Computation, error)
	Delete(id int64) (models.ComputationRead, error)

	}


