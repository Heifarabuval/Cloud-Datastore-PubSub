package models

type Webhook struct {
	ID     int64    `json:"id" validate:"required"`
	Fields []string `json:"fields" validate:"required"`
	Op     string   `json:"operator" validate:"required"`
}

type WebhookDto struct {
	Fields [] string `json:"fields" validate:"required"`
	Op     string   `json:"operator" validate:"required"`
}




