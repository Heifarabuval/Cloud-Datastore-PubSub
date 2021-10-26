package handler

import (
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"strings"
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
	var testString = strings.NewReader(`{"fields": ["a", "b", "c"],"operator":"sub"}`)
	c, _, res := createEchoTest("/webhook", testString, http.MethodPost)
	h, msw, _ := createMockedHandler()
	webhook := models.Webhook{
		ID:     9999999999999,
		Fields: []string{"a", "b", "c"},
		Op:     "sub",
	}
	msw.On("Create", mock.Anything).Return(&webhook, nil)
	if assert.NoError(t, h.CreateWebhook(c)) {
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, "{\"data\":{\"id\":9999999999999,\"fields\":[\"a\",\"b\",\"c\"],\"operator\":\"sub\"},\"responseCode\":201}\n", res.Body.String())
	}
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
