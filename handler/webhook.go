package handler

import (
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/datastoreHandlers"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/datastoreHandlers/webhookDatastore"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Json         interface{} `json:"data"`
	ResponseCode int         `json:"responseCode"`
}

func CreateWebhook(e *echo.Echo) {
	e.POST("/webhook", func(c echo.Context) (err error) {

		//Validate data
		dto := new(models.WebhookDto)

		if err = c.Bind(dto); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		if err = c.Validate(dto); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		dsHandler := webhookDatastore.InitClient(datastoreHandlers.CreateClient(context.Background()))

		//Persist data
		id, err := dsHandler.Create(dto.Op, dto.Fields)
		if err != nil {
			return echo.NewHTTPError(http.StatusConflict)
		}

		//Hydrate to return created object
		w := models.Webhook{
			ID:     id,
			Fields: dto.Fields,
			Op:     dto.Op,
		}
		response := Response{w, http.StatusCreated}

		return c.JSON(response.ResponseCode, response)
	})
}

func ReadAllWebhooks(e *echo.Echo) {
	e.GET("/webhook-all", func(c echo.Context) error {
		dsHandler := webhookDatastore.InitClient(datastoreHandlers.CreateClient(context.Background()))
		webhooks := dsHandler.ReadAll()
		response := Response{webhooks, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})
}

func ReadWebhook(e *echo.Echo) {
	e.GET("/webhook/:id", func(c echo.Context) error {

		_, id := datastoreHandlers.GetAndValidateId(c)

		dsHandler := webhookDatastore.InitClient(datastoreHandlers.CreateClient(context.Background()))

		entity, err := dsHandler.Read(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		response := Response{entity, http.StatusOK}
		return c.JSON(response.ResponseCode, response.Json)
	})
}

type WebhookDtoUpdate struct {
	Fields []string `json:"fields" validate:"min=2,dive,required" `
	Op     string   `json:"operator" validate:"eq=add|eq=sub|eq="`
}

func UpdateWebhook(e *echo.Echo) {
	e.PUT("/webhook/:id", func(c echo.Context) error {
		_, id := datastoreHandlers.GetAndValidateId(c)

		dto := new(WebhookDtoUpdate)

		if err := c.Bind(dto); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		if err := c.Validate(dto); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		dsHandler := webhookDatastore.InitClient(datastoreHandlers.CreateClient(context.Background()))

		//Persist data
		webhook, err := dsHandler.Update(id, dto.Op, dto.Fields)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		response := Response{webhook, http.StatusCreated}

		return c.JSON(response.ResponseCode, response)
	})
}

func DeleteWebhook(e *echo.Echo) {
	e.DELETE("/webhook/:id", func(c echo.Context) error {
		_, id := datastoreHandlers.GetAndValidateId(c)
		dsHandler := webhookDatastore.InitClient(datastoreHandlers.CreateClient(context.Background()))
		webhook, err := dsHandler.Delete(id)

		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		response := Response{webhook, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})
}

type Result struct {
	computation interface{}
	webhook     interface{}
}
