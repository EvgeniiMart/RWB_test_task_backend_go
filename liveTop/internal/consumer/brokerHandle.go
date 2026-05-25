package consumer

import (
	"encoding/json"
	"log"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
	"github.com/nats-io/nats.go"
)

func BrokerHandle(eventQueueWrap *joint.EventQueueWrapped,
	queriesMapWrap *joint.QueriesMapWrapped) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	_, err = nc.Subscribe("events", func(msg *nats.Msg) {
		var raw interface{}

		err := json.Unmarshal(msg.Data, &raw)
		if err != nil {
			log.Println("Invalid json:", err)
			return
		}

		err = validateJSON("event_schema.json", raw)
		if err != nil {
			log.Println("Schema validation failed:", err)
			return
		}

		var events []joint.Event

		err = json.Unmarshal(msg.Data, &events)
		if err != nil {
			log.Println("Parse failed:", err)
			return
		}

		storeEvents(eventQueueWrap, queriesMapWrap, events)
	})
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
