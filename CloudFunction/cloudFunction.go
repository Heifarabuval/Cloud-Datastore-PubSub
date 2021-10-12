package p

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	projectID    = "heifara-test"
	pubsubClient *pubsub.Client
	computeTopic *pubsub.Topic
)

func init() {
	log.Println("Init func")
	var psError error
	ctx := context.Background()
	pubsubClient, psError = pubsub.NewClient(ctx, projectID)

	if psError != nil {
		log.Fatal(psError)
	}

	computeTopic = pubsubClient.Topic("getCompute")
}

type Event struct {
	Subscription string `json:"subscription"`
	Message      struct {
		Attributes  map[string]string `json:"attributes"`
		Data        string            `json:"data"`
		PublishTime time.Time         `json:"publish_time"`
	} `json:"message"`
}

type PubSubPayload struct {
	WebhookId     int64       `json:"webhook_id"`
	ComputationId int64       `json:"computation_id"`
	Op            string      `json:"op"`
	Fields        []string    `json:"fields"`
	Values        []CustomMap `json:"values"`
	Result        int64       `json:"result"`
}

type CustomMap struct {
	Key   string `datastore:"key" json:"key"`
	Value int64  `datastore:"value" json:"value"`
}

func Main(w http.ResponseWriter, r *http.Request) {

	log.Println("Start the function")
	var psEvent Event
	var psPayload PubSubPayload

	if err := json.NewDecoder(r.Body).Decode(&psEvent); err != nil {
		log.Printf("json.NewDecoder error: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	data, err := base64.StdEncoding.DecodeString(psEvent.Message.Data)

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	if err := json.Unmarshal(data, &psPayload); err != nil {
		log.Println(err)
		return
	}

	psPayload.Result = compute(psPayload)

	ctx := context.Background()

	psPayloadJWT, _ := json.Marshal(&psPayload)

	if len(psPayloadJWT) == 0 {
		fmt.Fprint(w, "Bad request")
		return
	}

	_ = computeTopic.Publish(ctx, &pubsub.Message{
		Data: psPayloadJWT,
	})

	return

}

func compute(payload PubSubPayload) int64 {
	for key, value := range payload.Values {
		for key2, field := range payload.Fields {
			if value.Key == field {
				payload.Fields[key2] = strconv.FormatInt(value.Value, 10)
				log.Println("key1:", key, " key2:", key2, " value key", value.Key, " value Value ", value.Value, " field:", field)
			}
		}

	}
	finalCompute := payload.Fields

	result := payload.Result

	switch payload.Op {

	case "add":
		for _, s := range finalCompute {
			t, _ := strconv.Atoi(s)
			result = add(result, int64(t))
		}
	case "sub":
		for i, s := range finalCompute {
			t, _ := strconv.Atoi(s)
			if i == 0 {
				result = add(result, int64(t))
			} else {
				result = sub(result, int64(t))
			}

		}
	}
	log.Println("=====================================================")
	log.Println(result)
	log.Println("=====================================================")
	return result
}

func add(a int64, b int64) int64 {
	return a + b
}

func sub(a int64, b int64) int64 {
	return a - b
}
