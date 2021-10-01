package handler

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Response struct {
	Json         interface{} `json:"data"`
	ResponseCode int         `json:"responseCode"`
}

var (
	webhooks       = make(map[int64]models.Webhook)
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
		if dto.Op!="sub"&&dto.Op!="add" {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		//Persist data
		id := datastoreHandlers.Create(dto.Op,dto.Fields)

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
		webhooks:= datastoreHandlers.ReadAll()
		var response = Response{webhooks, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})
}

func ReadWebhook(e *echo.Echo) {
	e.GET("/webhook/:id", func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		//Validate data
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		datastoreHandlers.Read(id)

		var response = Response{id, http.StatusOK}
		return c.JSON(response.ResponseCode, response.Json)
	})
}

func UpdateWebhook(e *echo.Echo) {
	e.PUT("/webhook/:id", func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		//Check if id exist
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		//Check if webhook exist
		data, exist := webhooks[id]
		if exist != true {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		//Validate data
		dto := new(models.WebhookDto)
		if err := c.Bind(dto); err != nil {
			return err
		}
		if err := c.Validate(dto); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		if dto.Op != "" {
			data.Op = dto.Op
		}

		if len(dto.Fields) != 0 {
			data.Fields = dto.Fields
		}

		//Persist data
		webhooks[id] = data

		var response = Response{data, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})
}

func DeleteWebhook(e *echo.Echo) {
	e.DELETE("/webhook/:id", func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		//Validate data
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		value, exist := webhooks[id]
		if exist != true {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		//Delete & Persist data
		delete(webhooks, id)

		var response = Response{value, http.StatusOK}
		return c.JSON(response.ResponseCode, response)
	})
}
