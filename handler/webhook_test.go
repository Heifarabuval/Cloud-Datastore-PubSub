package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type WebhookDto struct {
	Fields []string `json:"fields" validate:"required,min=2,dive,required"`
	Op     string   `json:"operator" validate:"required,eq=add|eq=sub"`
}

func (m WebhookDto) InitWebhook(op string, fields []string) bool {

	m.Op = op
	m.Fields = fields
	return true
}

func TestCreateWebhook(t *testing.T) {
	// assert equality

	assert.Equal(t, 123, 123, "they should be equal")
}

func TestDeleteWebhook(t *testing.T) {
	// assert equality

	assert.Equal(t, 123, 123, "they should be equal")
}
func TestUpdateWebhook(t *testing.T) {
	// assert equality

	assert.Equal(t, 123, 123, "they should be equal")
}
func TestReadWebhook(t *testing.T) {

	assert.Equal(t, 123, 123, "they should be equal")
}
func TestReadAllWebhooks(t *testing.T) {
	// assert equality

	assert.Equal(t, 123, 123, "they should be equal")
}
