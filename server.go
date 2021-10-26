package main

import (
	"fmt"
	datastoreHandlers "github.com/Heifarabuval/Cloud-Datastore-PubSub/handler"
	"github.com/Heifarabuval/Cloud-Datastore-PubSub/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	wStore datastoreHandlers.StoreWebhook
	cStore datastoreHandlers.StoreComputation
)

//Init all routes
func initHandlers(e *echo.Echo, wStore datastoreHandlers.StoreWebhook, cStore datastoreHandlers.StoreComputation) {
	handler := datastoreHandlers.Handler{
		StoreWebhook:     wStore,
		StoreComputation: cStore,
	}

	//Crud webhookDatastore handler
	handler.AddCreateWebhook(e)
	handler.AddReadWebhook(e)
	handler.AddReadAllWebhooks(e)
	handler.AddUpdateWebhook(e)
	handler.AddDeleteWebhook(e)

	//Crud computationDatastore handler
	handler.AddCreateComputation(e)
	handler.AddReadAllComputations(e)
	handler.AddReadComputation(e)
	handler.AddDeleteComputation(e)
}

// Init a new client
func init() {
	ws, _ := datastoreHandlers.NewDatastoreWebhookStore(utils.DatastoreClient)
	wStore = ws

	cs, _ := datastoreHandlers.NewDatastoreComputationStore(utils.DatastoreClient)
	cStore = cs

}

func main() {
	//instantiate the web server
	e := echo.New()
	e.Validator = &datastoreHandlers.CustomValidator{Validator: validator.New()}

	//Getting port in .env
	port := "8000" /*utils.GetEnvVar("PORT", "8000")*/

	initHandlers(e, wStore, cStore)

	fmt.Printf("Server run on http://localhost:%s", port)

	//Run the webserver
	e.Logger.Fatal(e.Start(":" + port))

}
