package webhook

import (
	"Calicut/datastoreHandlers"
	"Calicut/models"
	"cloud.google.com/go/datastore"
	"context"
)

func Update(id int64,op string,fields[]string) interface{} {
	ctx := context.Background()
	client:= datastoreHandlers.CreateClient(ctx)

	webhook := &models.Webhook{}
	exist := datastoreHandlers.ReadById(id, "Webhook", webhook)
	if exist == nil {
		return exist
	}
	if len(op)==0 {
		op=webhook.Op
	}
	if len(fields)==0 {
		fields=webhook.Fields
	}


	key:=&datastore.Key{
		Kind:      "Webhook",
		ID:        id,
		Name:      "",
		Parent:    nil,
		Namespace: "",
	}

	entity, err := client.Put(ctx, key,
		&models.WebhookDto{
			Fields: fields,
			Op:     op,
		})

	if err != nil {
		return nil
	}

	return entity
}