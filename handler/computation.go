package handler

import (
	"Calicut/datastoreHandlers"
	"Calicut/datastoreHandlers/computationDatastore"
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
	ComputationId int64            `json:"computation_id"`
	Op            string           `json:"op"`
	Fields        []string         `json:"fields"`
	Values        []models.CustomMap `json:"values"`
	Result        int64            `json:"result"`
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

		//Persist data
		id := computationDatastore.Create(dto.WebhookId, dto.Values)

		if id == 0 {
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
		w := models.ComputationDto{
			WebhookId: dto.WebhookId,
			Result:    dto.Result,
			Values:    valueToStore,
			Computed:  dto.Computed,
		}

		entity := &models.Webhook{}
		res:= datastoreHandlers.ReadById(dto.WebhookId, "Webhook", entity)

		if res ==nil {
			return echo.NewHTTPError(http.StatusBadRequest)

		}

		pl:= PubSubPayload{
			ComputationId: id,
			Op:            entity.Op,
			Fields:        entity.Fields,
			Values:        w.Values,
			Result:        0,
		}

		psPayload, _ := json.Marshal(pl)
		ctx := context.Background()

		_ = computeTopic.Publish(ctx, &pubsub.Message{
			Data: psPayload,
		})


		var response = Response{w, http.StatusCreated}

		if err != nil {
			return err
		}

		return c.JSON(response.ResponseCode, response)
	})

}

func ReadComputation(e *echo.Echo) {
	e.GET("/computation/:id", func(c echo.Context) error {

		_, id := datastoreHandlers.GetAndValidateId(c)
		/*	formValue := c.FormValue("webhook")
			if formValue == "1" {
				datastoreHandlers.ReadComputationByWebhookId(id)
			}*/

		entity := computationDatastore.Read(id)
		if entity == nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		var response = Response{entity, http.StatusOK}
		return c.JSON(response.ResponseCode, response.Json)
	})
}

func ReadAllComputation(e *echo.Echo) {
	e.GET("/computation-all", func(c echo.Context) error {
		computations := computationDatastore.ReadAll()

		var response = Response{computations, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})
}

func UpdateComputation() {

}

func DeleteComputation(e *echo.Echo) {
	e.DELETE("/computation/:id", func(c echo.Context) error {
		_, id := datastoreHandlers.GetAndValidateId(c)
		entity := computationDatastore.Delete(id)
		if entity == nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		var response = Response{entity, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})

}
