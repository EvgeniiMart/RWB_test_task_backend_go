package consumer

import (
	"log"
	"time"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
)

func storeEvents(eventQueueWrap *joint.EventQueueWrapped,
	queriesMapWrap *joint.QueriesMapWrapped, events []joint.Event) {
	eventQueueWrap.Mu.Lock()
	queriesMapWrap.Mu.Lock()
	defer eventQueueWrap.Mu.Unlock()
	defer queriesMapWrap.Mu.Unlock()

	for _, event := range events {
		if event.Timestamp.IsZero() {
			event.Timestamp = time.Now()
		}

		log.Printf("Received event: query=%s, delta=%d, timestamp=%s\n",
			event.Query, event.Delta, event.Timestamp.Format(time.RFC3339))

		queriesMapWrap.Data[event.Query] += event.Delta

		eventQueueWrap.Data.PushBack(event)
	}
}
