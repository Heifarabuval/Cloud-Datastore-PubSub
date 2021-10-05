package datastoreHandlers

import (
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func ReadById(id int64, kind string, model interface{}) interface{} {


	ctx := context.Background()
	client := CreateClient(ctx)

	//Create key for search
	key := &datastore.Key{
		Kind:      kind,
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}

	entity := model

	err := client.Get(ctx, key, entity)

	if err != nil {
		return nil
	}

	defer client.Close()

	return entity

}

func ReadComputationByWebhookId(id int64) models.ComputationRead {
	ctx := context.Background()
	client := CreateClient(ctx)
	var entity  []models.ComputationRead
	query := datastore.NewQuery("Computation").Filter("webhookId =", id)
	_,err := client.GetAll(ctx,query,&entity)

	if err != nil{
		return models.ComputationRead{}
	}
	return entity[0]
}


func CreateClient(ctx context.Context) *datastore.Client {
	//Create client
	projectID := "heifara-test"
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

func GetAndValidateId(c echo.Context) (*echo.HTTPError, int64) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || len(c.Param("id"))==0 {
		return echo.NewHTTPError(http.StatusBadRequest), 0
	}
	return nil, id
}
