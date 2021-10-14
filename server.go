package main

import (
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/handler"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	//instantiate the web server
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	//Getting port in .env
	port := utils.GetEnvVar("PORT", "8000")

	//Crud webhookDatastore handler
	handler.CreateWebhook(e)
	handler.ReadWebhook(e)
	handler.ReadAllWebhooks(e)
	handler.UpdateWebhook(e)
	handler.DeleteWebhook(e)

	//Crud computationDatastore handler
	handler.CreateComputation(e)
	handler.ReadAllComputation(e)
	handler.ReadComputation(e)
	handler.DeleteComputation(e)

	fmt.Printf("Server run on http://localhost:%s", port)

	//Run the webserver
	e.Logger.Fatal(e.Start(":" + port))

}
