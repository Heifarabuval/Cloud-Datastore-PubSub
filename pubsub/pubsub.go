package pubsub

import (
	"Calicut/config"
	"cloud.google.com/go/pubsub"
	"context"
)

var ctx= context.Background()
var projectID, projectIdError = config.GetEnvConst("PROJECT_NAME")
var pubsubClient, psError = pubsub.NewClient(ctx, projectID)
var computeTopic = pubsubClient.Topic("compute")

func initPs()  {

}

