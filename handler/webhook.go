package handler

import (
	"Calicut/datastoreHandlers"
	"Calicut/datastoreHandlers/webhookDatastore"
	"Calicut/models"
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

		//Persist data
		id := webhookDatastore.Create(dto.Op, dto.Fields)
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
		webhooks := webhookDatastore.ReadAll()
		var response = Response{webhooks, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})
}

func ReadWebhook(e *echo.Echo) {
	e.GET("/webhook/:id", func(c echo.Context) error {

		_, id := datastoreHandlers.GetAndValidateId(c)

		entity := webhookDatastore.Read(id)
		if entity == nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		var response = Response{entity, http.StatusOK}
		return c.JSON(response.ResponseCode, response.Json)
	})
}




type WebhookDtoUpdate struct {
	Fields []string `json:"fields" validate:"dive,required"`
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

		//Persist data
		key := webhookDatastore.Update(id, dto.Op, dto.Fields)

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

		entity := webhookDatastore.Delete(id)
		if entity == nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		var response = Response{entity, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})
}

type Result struct {
	computation interface{}
	webhook interface{}
}

