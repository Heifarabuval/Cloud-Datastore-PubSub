package handler

import (
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/handler/datastoreHandlers"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Json         interface{} `json:"data"`
	ResponseCode int         `json:"responseCode"`
}

/*===========================================     CREATE       =========================================================*/

func (h *Handler) AddCreateWebhook(e *echo.Echo) {
	e.POST("/webhook", h.CreateWebhook)
}

func (h *Handler) CreateWebhook(c echo.Context) (err error) {

	//Validate data
	dto := new(models.WebhookDto)

	if err = c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err = c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	payload := models.Webhook{
		ID:     0,
		Fields: dto.Fields,
		Op:     dto.Op,
	}

	//Persist data
	w, err := h.StoreWebhook.Create(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict)
	}

	response := Response{w, http.StatusCreated}

	return c.JSON(response.ResponseCode, response)

}

/*===========================================     READ ALL     =========================================================*/

func (h *Handler) AddReadAllWebhooks(e *echo.Echo) {
	e.GET("/webhook-all", h.ReadAllWebhooks)
}
func (h *Handler) ReadAllWebhooks(c echo.Context) (err error) {
	w, err := h.StoreWebhook.ReadAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict)
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusConflict)
	}
	response := Response{w, http.StatusOK}
	return c.JSON(response.ResponseCode, response)

}

/*=============================================     READ      ==========================================================*/

func (h *Handler) AddReadWebhook(e *echo.Echo) {
	e.GET("/webhook/:id", h.ReadWebhook)
}
func (h *Handler) ReadWebhook(c echo.Context) (err error) {

	_, id := datastoreHandlers.GetAndValidateId(c)

	w, err := h.StoreWebhook.Read(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict)
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	response := Response{w, http.StatusOK}
	return c.JSON(response.ResponseCode, response.Json)

}

type WebhookDtoUpdate struct {
	Fields []string `json:"fields" validate:"min=2,dive,required" `
	Op     string   `json:"operator" validate:"eq=add|eq=sub|eq="`
}

/*===========================================      UPDATE      =========================================================*/

func (h Handler) AddUpdateWebhook(e *echo.Echo) {
	e.PUT("/webhook/:id", h.UpdateWebhook)
}
func (h *Handler) UpdateWebhook(c echo.Context) (err error) {

	_, id := datastoreHandlers.GetAndValidateId(c)

	dto := new(WebhookDtoUpdate)

	if err := c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err := c.Validate(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}


	//Persist data
	webhook, err := h.StoreWebhook.Update(id, dto.Op, dto.Fields)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	response := Response{webhook, http.StatusCreated}

	return c.JSON(response.ResponseCode, response)
}

/*===========================================      DELETE      =========================================================*/

func (h Handler) AddDeleteWebhook(e *echo.Echo) {
	e.DELETE("/webhook/:id", h.DeleteWebhook)
}

func (h Handler) DeleteWebhook(c echo.Context) (err error) {
	_, id := datastoreHandlers.GetAndValidateId(c)
	webhook, err := h.StoreWebhook.Delete(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	response := Response{webhook, http.StatusOK}
	return c.JSON(response.ResponseCode, response)
}

type Result struct {
	computation interface{}
	 webhook     interface{}
}