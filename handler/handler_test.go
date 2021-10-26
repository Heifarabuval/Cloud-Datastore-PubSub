package handler

import (
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
)

type (
	MockedStoreWebhook struct {
		mock.Mock
		StoreWebhook
	}

	MockedStoreComputation struct {
		mock.Mock
		StoreComputation
	}
)

func (s *MockedStoreWebhook) Read(id int64) (models.Webhook, error) {
	args := s.Called(id)
	return args.Get(0).(models.Webhook), args.Error(1)
}

func (s *MockedStoreWebhook) ReadAll() ([]models.Webhook, error) {
	args := s.Called()
	return args.Get(0).([]models.Webhook), args.Error(1)
}

func (s *MockedStoreWebhook) Create(webhook *models.Webhook) (*models.Webhook, error) {
	args := s.Called(webhook)
	webhook.ID = 333333333333
	return args.Get(0).(*models.Webhook), args.Error(1)
}

func (s *MockedStoreWebhook) Update(id int64, op string, fields []string) (models.Webhook, error) {
	webhook := models.Webhook{
		ID:     345678945678,
		Fields: fields,
		Op:     op,
	}
	args := s.Called(webhook)
	return args.Get(0).(models.Webhook), args.Error(1)
}

func (s *MockedStoreWebhook) Delete(id int64) (models.Webhook, error) {
	args := s.Called(id)
	return args.Get(0).(models.Webhook), args.Error(1)
}

func createMockedHandler() (h Handler, mockedStoreWebhook *MockedStoreWebhook, mockedStoreComputation *MockedStoreComputation) {
	mockedStoreWebhook = &MockedStoreWebhook{}
	mockedStoreComputation = &MockedStoreComputation{}
	var storeWebhook StoreWebhook = mockedStoreWebhook
	var storeComputation StoreComputation = mockedStoreComputation
	h = Handler{
		StoreWebhook:     storeWebhook,
		StoreComputation: storeComputation,
	}
	return
}

func createEchoTest(path string, reader io.Reader, method string) (c echo.Context, req *http.Request, res *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	req = httptest.NewRequest(method, path, reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res = httptest.NewRecorder()
	c = e.NewContext(req, res)
	c.SetPath(path)
	return
}
