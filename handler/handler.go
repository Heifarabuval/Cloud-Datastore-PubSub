package datastoreHandlers

import (
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
)

type StoreComputation interface {
	Create(webhookId int64, values map[string]int64) (int64, error)
	Read(id int64) (models.Computation, error)
	ReadAll() ([]models.Computation, error)
	Delete(id int64) (models.Computation, error)
}

type StoreWebhook interface {
	Create(op string, fields []string) (int64, error)
	Read(id int64) (models.Webhook, error)
	ReadAll() []models.Webhook
	Update(id int64, op string, fields []string) (models.Webhook, error)
	Delete(id int64) (models.Webhook, error)
}

type Handler struct {
	StoreWebhook     StoreWebhook
	StoreComputation StoreComputation
}
