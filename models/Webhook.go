package models

import (
	"cloud.google.com/go/datastore"
)


type Webhook struct {
	ID     int64    `json:"id" validate:"required" `
	Fields []string `json:"fields" validate:"required,isFieldsAreNotBlanks"`
	Op     string   `json:"operator" validate:"required"`
}

type WebhookDto struct {
	Fields [] string `json:"fields" validate:"required,min=2"`
	Op     string   `json:"operator" validate:"required,eq=add|eq=sub"`
}

type WebhookDtoUpdate struct {
	Fields [] string `json:"fields"`
	Op     string   `json:"operator" validate:"eq=add|eq=sub|eq="`
}



func (w *Webhook) LoadKey(k *datastore.Key) error {
	w.ID = k.ID
	return nil
}

func (w *Webhook) Load(ps []datastore.Property) error {
	return datastore.LoadStruct(w, ps)
}

func (w *Webhook) Save() ([]datastore.Property, error) {
	return datastore.SaveStruct(w)
}

