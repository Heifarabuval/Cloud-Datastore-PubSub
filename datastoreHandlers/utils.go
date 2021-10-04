package datastoreHandlers

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func ReadById(id int64, kind string, model interface{}) interface{} {

	//Creating client
	//TODO ENCAPSULATE
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
