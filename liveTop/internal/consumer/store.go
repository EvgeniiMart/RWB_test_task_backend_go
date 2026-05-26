package consumer

import (
	"time"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
)

// Logic of work with shared data storages:
// 1. delta is taken into account for queriesMap
// 2. whole event is added to eventQueue for future deletion (after 5 minutes)
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

		queriesMapWrap.Data[event.Query] += event.Delta

		eventQueueWrap.Data.PushBack(event)
	}
}
