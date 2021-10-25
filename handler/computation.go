package datastoreHandlers

import (
	"context"
	"encoding/json"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/handler/datastoreHandlers"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/handler/datastoreHandlers/computationDatastore"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/handler/datastoreHandlers/webhookDatastore"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/utils"
	"github.com/labstack/echo/v4"
)

type ComputationDtoCreate struct {
	WebhookId int64            `json:"webhookId" json:"webhookId" validate:"required"`
	Values    map[string]int64 `json:"values"`
	Result    int64            `json:"result"`
	Computed  bool             `json:"computed"`
}

var (
	projectID    = utils.GetEnvVar("PROJECT_NAME", "heifara-test")
	pubSubClient *pubsub.Client
	computeTopic *pubsub.Topic
)

func init() {
	var psError error
	ctx := context.Background()
	pubSubClient, psError = pubsub.NewClient(ctx, projectID)

	if psError != nil {
		log.Fatal(psError)
	}
	computeTopic = pubSubClient.Topic("compute")
}

type PubSubPayload struct {
	WebhookId     int64              `json:"webhook_id"`
	ComputationId int64              `json:"computation_id"`
	Op            string             `json:"op"`
	Fields        []string           `json:"fields"`
	Values        []models.CustomMap `json:"values"`
	Result        int64              `json:"result"`
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

/*===========================================     CREATE       =========================================================*/

func (h *Handler) AddCreateComputation(e *echo.Echo) {
	e.POST("/computation", h.CreateComputation)
}
func (h *Handler) CreateComputation(c echo.Context) (err error) {

	//Validate data
	dto := new(ComputationDtoCreate)
	if err = c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err = c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Looking for linked webhook
	dsHandlerWebhook := webhookDatastore.InitClient(datastoreHandlers.CreateClient(context.Background()))
	webhook, err := dsHandlerWebhook.Read(dto.WebhookId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	//Check if fields and values concords
	if len(webhook.Fields) != len(dto.Values) {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	for s := range dto.Values {
		if contains(webhook.Fields, s) == false {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
	}

	//Persist data
	dsHandlerComputation := computationDatastore.InitClient(datastoreHandlers.CreateClient(context.Background()))
	id, err := dsHandlerComputation.Create(dto.WebhookId, dto.Values)

	if err != nil {
		return echo.NewHTTPError(http.StatusConflict)
	}

	var valueToStore []models.CustomMap

	for key, value := range dto.Values {
		item := models.CustomMap{
			Key:   key,
			Value: value,
		}
		valueToStore = append(valueToStore, item)
	}

	//Hydrate to return created object
	w := models.Computation{
		ID:        id,
		WebhookId: dto.WebhookId,
		Result:    dto.Result,
		Values:    dto.Values,
		Computed:  dto.Computed,
	}

	pubSubPayload := PubSubPayload{
		WebhookId:     webhook.ID,
		ComputationId: id,
		Op:            webhook.Op,
		Fields:        webhook.Fields,
		Values:        valueToStore,
		Result:        0,
	}

	psPayload, _ := json.Marshal(pubSubPayload)
	ctx := context.Background()

	computeTopic.Publish(ctx, &pubsub.Message{
		Data: psPayload,
	})

	response := Response{w, http.StatusCreated}

	if err != nil {
		return err
	}

	return c.JSON(response.ResponseCode, response)

}

/*===========================================     READ ALL     =========================================================*/

func (h *Handler) AddReadAllComputations(e *echo.Echo) {
	e.GET("/computation-all", h.ReadAllComputations)
}

func (h *Handler) ReadAllComputations(c echo.Context) (err error) {
	dsHandler := computationDatastore.InitClient(datastoreHandlers.CreateClient(context.Background()))
	computations, err := dsHandler.ReadAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if len(computations) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Computations list is empty")
	}
	response := Response{computations, http.StatusOK}
	return c.JSON(response.ResponseCode, response)
}

/*=============================================     READ      ==========================================================*/

func (h *Handler) AddReadComputation(e *echo.Echo) {
	e.GET("/computation/:id", h.ReadComputation)
}
func (h *Handler) ReadComputation(c echo.Context) (err error) {

	_, id := datastoreHandlers.GetAndValidateId(c)

	dsHandler := computationDatastore.InitClient(datastoreHandlers.CreateClient(context.Background()))
	computation, err := dsHandler.Read(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	response := Response{computation, http.StatusOK}
	return c.JSON(response.ResponseCode, response.Json)

}

/*===========================================      DELETE      =========================================================*/

func (h *Handler) AddDeleteComputation(e *echo.Echo) {
	e.DELETE("/computation/:id", h.DeleteComputation)
}
func (h *Handler) DeleteComputation(c echo.Context) (err error) {
	_, id := datastoreHandlers.GetAndValidateId(c)
	dsHandler := computationDatastore.InitClient(datastoreHandlers.CreateClient(context.Background()))
	computation, err := dsHandler.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	response := Response{computation, http.StatusOK}
	return c.JSON(response.ResponseCode, response)

}
