package datastoreHandlers

import (
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"log"
	"strconv"
)
func Read(id int64)  {
	ctx := context.Background()

	projectID := "heifara-test"

	client, err := datastore.NewClient(ctx,projectID)

	if err != nil{
		log.Fatalf("Failed to create client: %v",err)
	}
	var webhooks []*models.Webhook
	key := datastore.NameKey("Article", strconv.FormatInt(id, 10), nil)
	find := client.Get(ctx,key,webhooks)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("test")

	fmt.Println(find)

	defer client.Close()


}

