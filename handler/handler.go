package handler

import "Calicut/models"

type StoreWebhook interface {
	Get(id int64) (webhook *models.Webhook, err error)
	GetAll() (webhooks []*models.Webhook)
}

