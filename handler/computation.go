package handler

import (
	"Calicut/datastoreHandlers"
	"Calicut/datastoreHandlers/computationDatastore"
	"Calicut/datastoreHandlers/webhookDatastore"
	"Calicut/models"
	"Calicut/utils"
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type ComputationDtoCreate struct {
	WebhookId int64            `json:"webhookId" json:"webhookId" validate:"required"`
	Values    map[string]int64 `json:"values"`
	Result    int64            `json:"result"`
	Computed  bool             `json:"computed"`
}

var (
	projectID    = utils.GetEnvVar("PROJECT_NAME", "heifara-test")
	pubsubClient *pubsub.Client
	computeTopic *pubsub.Topic
)

func init() {
	var psError error
	ctx := context.Background()
	pubsubClient, psError = pubsub.NewClient(ctx, projectID)

	if psError != nil {
		log.Fatal(psError)
	}
	computeTopic = pubsubClient.Topic("compute")
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

func CreateComputation(e *echo.Echo) {

	e.POST("/computation", func(c echo.Context) (err error) {

		//Validate data
		dto := new(ComputationDtoCreate)
		if err = c.Bind(dto); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err = c.Validate(dto); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		webhook, err := webhookDatastore.Read(dto.WebhookId)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		//Check if fields and values concords
		if len(webhook.Fields) != len(dto.Values) {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		for s, _ := range dto.Values {
			if contains(webhook.Fields, s) == false {
				return echo.NewHTTPError(http.StatusBadRequest)
			}
		}

		//Persist data
		id, err := computationDatastore.Create(dto.WebhookId, dto.Values)

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
	})

}

func ReadComputation(e *echo.Echo) {
	e.GET("/computation/:id", func(c echo.Context) error {
		_, id := datastoreHandlers.GetAndValidateId(c)
		entity, err := computationDatastore.Read(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		response := Response{entity, http.StatusOK}
		return c.JSON(response.ResponseCode, response.Json)
	})
}

func ReadAllComputation(e *echo.Echo) {
	e.GET("/computation-all", func(c echo.Context) error {
		computations, err := computationDatastore.ReadAll()
		println(len(computations))
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		if len(computations) == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "Computations list is empty")
		}
		response := Response{computations, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})
}

func UpdateComputation() {

}

func DeleteComputation(e *echo.Echo) {
	e.DELETE("/computation/:id", func(c echo.Context) error {
		_, id := datastoreHandlers.GetAndValidateId(c)
		computation,err := computationDatastore.Delete(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		response := Response{computation, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})

}
