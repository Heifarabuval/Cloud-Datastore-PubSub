package models

import "cloud.google.com/go/datastore"

type Computation struct {
	ID        int64            `json:"id" validate:"required"`
	WebhookId int64            `json:"webhookId" validate:"required"`
	Result    int64            `json:"result" validate:"required"`
	Values    map[string]int64 `json:"values" validate:"required"`
	Computed  bool             `json:"computed" validate:"required"`
}

type ComputationDto struct {
	WebhookId int64       `json:"webhookId" datastore:"webhookId" validate:"required"`
	Values    []CustomMap `json:"values" datastore:"values"`
	Result    int64       `json:"result" datastore:"result"`
	Computed  bool        `json:"computed" datastore:"computed"`
}

type ComputationRead struct {
	ID        int64              `json:"id" validate:"required"`
	WebhookId int64              `json:"webhookId" validate:"required"`
	Result    int64              `json:"result" validate:"required"`
	Values    []CustomMap `json:"values" validate:"required"`
	Computed  bool               `json:"computed" validate:"required"`
}

func (c ComputationRead) TransformToMap(computation []CustomMap) map[string]int64 {
	m := make(map[string]int64)
	for _, v := range computation {
		m[v.Key] = v.Value
	}
	return m
}

func (c *ComputationRead) LoadKey(k *datastore.Key) error {
	c.ID = k.ID
	return nil
}

func (c *ComputationRead) Load(ps []datastore.Property) error {
	return datastore.LoadStruct(c, ps)
}

func (c *ComputationRead) Save() ([]datastore.Property, error) {
	return datastore.SaveStruct(c)
}
