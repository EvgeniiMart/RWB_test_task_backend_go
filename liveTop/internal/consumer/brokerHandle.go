package consumer

import (
	"encoding/json"
	"log"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
	"github.com/nats-io/nats.go"
)

// All broker handling logic (separated from network code)
func processBrokerMessage(msg *nats.Msg, eventQueueWrap *joint.EventQueueWrapped,
	queriesMapWrap *joint.QueriesMapWrapped, cfg *joint.Config) {
	var raw interface{}

	err := json.Unmarshal(msg.Data, &raw)
	if err != nil {
		log.Println("Invalid json:", err)
		return
	}

	err = validateJSON("data/contract.schema.json", raw)
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

	storeEvents(eventQueueWrap, queriesMapWrap, events, cfg)
}

// NATS subscription for broker (separated from message processing logic)
func BrokerHandle(eventQueueWrap *joint.EventQueueWrapped,
	queriesMapWrap *joint.QueriesMapWrapped, cfg *joint.Config) {
	nc, err := nats.Connect(cfg.NATSurl)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	_, err = nc.Subscribe("events", func(msg *nats.Msg) {
		processBrokerMessage(msg, eventQueueWrap, queriesMapWrap, cfg)
	})
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
