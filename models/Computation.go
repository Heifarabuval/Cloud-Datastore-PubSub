package models

type Computation struct {
	Id        int64            `json:"id" validate:"required"`
	WebhookId int64            `json:"webhookId" validate:"required"`
	Result    int64            `json:"result" validate:"required"`
	Values    map[string]int64 `json:"values" validate:"required"`
}

type ComputationDto struct {
	WebhookId int64            `json:"webhookId" validate:"required"`
	Values    map[string]int64 `json:"values" validate:"required"`
}
