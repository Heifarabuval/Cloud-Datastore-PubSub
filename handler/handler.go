package handler

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

/*======================================= GLOBALS ==============================================*/

type (
	CustomValidator struct {
		Validator *validator.Validate
	}
)

type Handler struct {
	StoreWebhook     StoreWebhook
	StoreComputation StoreComputation
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

/*======================================= WEBHOOK HANDLERS ==============================================*/

type StoreWebhook interface {
	Create(webhook *models.Webhook) (*models.Webhook, error)
	Read(id int64) (models.Webhook, error)
	ReadAll() ([]models.Webhook, error)
	Update(id int64, op string, fields []string) (models.Webhook, error)
	Delete(id int64) (models.Webhook, error)
}

type DsWebhookStore struct {
	StoreWebhook
	client *datastore.Client
}

func NewDatastoreWebhookStore(client *datastore.Client) (StoreWebhook, error) {
	var ds StoreWebhook = &DsWebhookStore{
		client: client,
	}
	return ds, nil
}

func (s *DsWebhookStore) Create(webhook *models.Webhook) (*models.Webhook, error) {
	newKey := datastore.IncompleteKey("Webhook", nil)

	key, err := s.client.Put(context.Background(), newKey,
		webhook)

	_, err = s.client.Put(context.Background(), key,
		&models.Webhook{
			ID:     key.ID,
			Fields: webhook.Fields,
			Op:     webhook.Op,
		})

	w := models.Webhook{
		ID:     key.ID,
		Fields: webhook.Fields,
		Op:     webhook.Op,
	}

	if err != nil {
		return &w, err
	}

	return &w, nil
}

func (s *DsWebhookStore) Read(id int64) (models.Webhook, error) {
	webhook := models.Webhook{}

	//Create key for search
	key := &datastore.Key{
		Kind:      "Webhook",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}

	err := s.client.Get(context.Background(), key, &webhook)
	webhook.ID = key.ID
	if err != nil {
		return webhook, err
	}

	return webhook, nil
}

func (s *DsWebhookStore) ReadAll() ([]models.Webhook, error) {
	//Model
	var webhooks []models.Webhook

	//Request
	_, err := s.client.GetAll(context.Background(), datastore.NewQuery("Webhook"), &webhooks)

	if err != nil {
		return webhooks, err
	}

	return webhooks, nil
}

func (s *DsWebhookStore) Update(id int64, op string, fields []string) (models.Webhook, error) {
	ctx := context.Background()
	webhook, err := s.Read(id)
	if err != nil {
		return webhook, err
	}

	if len(op) == 0 {
		op = webhook.Op
	}
	if len(fields) == 0 {
		fields = webhook.Fields
	}

	key := &datastore.Key{
		Kind:      "Webhook",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}

	_, err = s.client.Put(ctx, key,
		&models.Webhook{
			ID:     id,
			Fields: fields,
			Op:     op,
		})

	webhook, err = s.Read(id)
	if err != nil {
		return webhook, nil
	}

	if err != nil {
		return webhook, nil
	}

	return webhook, nil
}
func (s *DsWebhookStore) Delete(id int64) (models.Webhook, error) {
	//Verify if webhookDatastore exist
	webhook, err := s.Read(id)
	if err != nil {
		return webhook, err
	}

	//Create key for search
	key := &datastore.Key{
		Kind:      "Webhook",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}

	_ = s.client.Delete(context.Background(), key)

	return webhook, nil
}

/*======================================= COMPUTATIONS HANDLERS ==============================================*/

type StoreComputation interface {
	Create(webhookId int64, values map[string]int64) (int64, error)
	Read(id int64) (models.Computation, error)
	ReadAll() ([]models.Computation, error)
	Delete(id int64) (models.Computation, error)
}

type DsComputationStore struct {
	StoreComputation
	client *datastore.Client
}

func NewDatastoreComputationStore(client *datastore.Client) (StoreComputation, error) {
	var ds StoreComputation = &DsComputationStore{
		client: client,
	}
	return ds, nil
}


func (s *DsComputationStore) Create(webhookId int64, values map[string]int64) (int64, error){
	//Transform map for pub/sub
	var valueToStore []models.CustomMap
	for key, value := range values {
		item := models.CustomMap{
			Key:   key,
			Value: value,
		}
		valueToStore = append(valueToStore, item)
	}

	newKey := datastore.IncompleteKey("Computation", nil)

	computation, err := s.client.Put(context.Background(), newKey,
		&models.ComputationDto{
			WebhookId: webhookId,
			Values:    valueToStore,
			Result:    0,
			Computed:  false,
		})

	if err != nil {
		return 0, err
	}

	return computation.ID, nil
}
func (s *DsComputationStore) Read(id int64) (models.Computation, error){
	computationDs := models.ComputationRead{}

	//Create key for search
	key := &datastore.Key{
		Kind:      "Computation",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}

	err := s.client.Get(context.Background(), key, &computationDs)
	computation := models.Computation{}

	if err != nil {
		return computation, err
	}

	//Hydrate data
	computation.ID = computationDs.ID
	computation.WebhookId = computationDs.WebhookId
	computation.Result = computationDs.Result
	computation.Values = computationDs.TransformToMap(computationDs.Values)
	computation.Computed = computationDs.Computed
	return computation, nil
}
func (s *DsComputationStore) ReadAll() ([]models.Computation, error){
	//Model
	var computations []models.ComputationRead
	var computationsFinal []models.Computation

	//Request
	_, err := s.client.GetAll(context.Background(), datastore.NewQuery("Computation"), &computations)

	for i, v := range computations {

		computation := models.Computation{
			ID:        v.ID,
			WebhookId: v.WebhookId,
			Result:    v.Result,
			Values:    computations[i].TransformToMap(computations[i].Values),
			Computed:  v.Computed,
		}
		computationsFinal = append(computationsFinal, computation)
	}

	if err != nil {
		return []models.Computation{}, err
	}

	return computationsFinal, nil
}
func (s *DsComputationStore) Delete(id int64) (models.Computation, error){
	//Verify if computationDatastore exist
	ctx := context.Background()
	computation, err := s.Read(id)

	if err != nil {
		return computation, err
	}

	key := &datastore.Key{
		Kind:      "Computation",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}
	_ = s.client.Delete(ctx, key)

	return computation, nil
}

