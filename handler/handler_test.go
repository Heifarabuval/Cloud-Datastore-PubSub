package datastoreHandlers

import (
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
	"github.com/stretchr/testify/mock"
)

type (
	MockedStoreWebhook struct {
		StoreWebhook
		mock.Mock
	}

	MockedStoreComputation struct {
		StoreComputation
		mock.Mock
	}
)

func (s *MockedStoreWebhook) Read(id int64) (*models.Webhook, error) {
	args := s.Called(id)
	return args.Get(0).(*models.Webhook), args.Error(1)
}

func (s *MockedStoreWebhook) ReadAll(id int64) ([]*models.Webhook, error) {
	args := s.Called()
	return args.Get(0).([]*models.Webhook), args.Error(1)
}

func CreateMockedHandler() (h Handler, mockedStoreWebhook *MockedStoreWebhook, mockedStoreComputation *MockedStoreComputation) {
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

/*func CreateEchoTest(path string, reader io.Reader, method string)(c echo.Context, req *http.Request, res *httptest.ResponseRecorder){
	e:= echo.New()

}*/

