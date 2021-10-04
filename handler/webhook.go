package handler

import (
	"Calicut/datastoreHandlers"
	"Calicut/datastoreHandlers/webhook"
	"Calicut/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Json         interface{} `json:"data"`
	ResponseCode int         `json:"responseCode"`
}

var (
	webhooks = make(map[int64]models.Webhook)
)

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

		//Validate fields notblank
		for i := 0; i < len(dto.Fields); i++ {
			if len(dto.Fields[i]) == 0 {
				return echo.NewHTTPError(http.StatusBadRequest)
			}
		}

		//Persist data
		id := webhook.Create(dto.Op, dto.Fields)
		if id == 0 {
			return echo.NewHTTPError(http.StatusConflict)
		}

		//Hydrate to return created object
		w := models.Webhook{
			ID:     id,
			Fields: dto.Fields,
			Op:     dto.Op,
		}
		var response = Response{w, http.StatusCreated}

		return c.JSON(response.ResponseCode, response)
	})
}

func ReadAllWebhooks(e *echo.Echo) {
	e.GET("/webhook-all", func(c echo.Context) error {
		webhooks := webhook.ReadAll()
		var response = Response{webhooks, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})
}

func ReadWebhook(e *echo.Echo) {
	e.GET("/webhook/:id", func(c echo.Context) error {

		_, id := datastoreHandlers.GetAndValidateId(c)

		entity := webhook.Read(id)
		if entity == nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		var response = Response{entity, http.StatusOK}
		return c.JSON(response.ResponseCode, response.Json)
	})
}

func UpdateWebhook(e *echo.Echo) {
	e.PUT("/webhook/:id", func(c echo.Context) error {
		_, id := datastoreHandlers.GetAndValidateId(c)

		dto := new(models.WebhookDtoUpdate)

		if err := c.Bind(dto); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		if err := c.Validate(dto); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		//Validate fields notblank
		for i := 0; i < len(dto.Fields); i++ {
			if len(dto.Fields[i]) == 0 {
				return echo.NewHTTPError(http.StatusBadRequest)
			}
		}

		//Persist data
		key := webhook.Update(id, dto.Op, dto.Fields)

		if key == nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		entity := &models.Webhook{}
		datastoreHandlers.ReadById(id, "Webhook", entity)
		var response = Response{entity, http.StatusCreated}

		return c.JSON(response.ResponseCode, response)
	})
}

func DeleteWebhook(e *echo.Echo) {
	e.DELETE("/webhook/:id", func(c echo.Context) error {
		_, id := datastoreHandlers.GetAndValidateId(c)

		entity := webhook.Delete(id)
		if entity == nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		var response = Response{entity, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})
}
