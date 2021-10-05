package handler

import (
	"Calicut/datastoreHandlers"
	"Calicut/datastoreHandlers/computationDatastore"
	"Calicut/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ComputationDtoCreate struct {
	WebhookId int64            `json:"webhookId" json:"webhookId" validate:"required"`
	Values    map[string]int64 `json:"values" json:"webhookId"`
	Result    int64            `json:"result" json:"webhookId"`
	Computed  bool             `json:"computed" json:"webhookId"`
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

		//Hydrate to return created object
		w := models.Computation{
			ID:        id,
			WebhookId: dto.WebhookId,
			Result:    dto.Result,
			Values:    dto.Values,
			Computed:  dto.Computed,
		}
		var response = Response{w, http.StatusCreated}

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
