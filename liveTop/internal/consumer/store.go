package consumer

import (
	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
)

func storeEvents(eventQueueWrap *joint.EventQueueWrapped,
	queriesMapWrap *joint.QueriesMapWrapped, events []joint.Event) {
	eventQueueWrap.Mu.Lock()
	queriesMapWrap.Mu.Lock()
	defer eventQueueWrap.Mu.Unlock()
	defer queriesMapWrap.Mu.Unlock()

	for _, event := range events {
		queriesMapWrap.Data[event.Query] += event.Delta

		eventQueueWrap.Data.PushBack(event)
	}
}
