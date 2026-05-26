package looper

import (
	"log"
	"sort"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"

	"time"
)

func processEvents(eventQueueWrap *joint.EventQueueWrapped,
	queriesMapWrap *joint.QueriesMapWrapped) {
	eventQueueWrap.Mu.Lock()
	defer eventQueueWrap.Mu.Unlock()

	for eventQueueWrap.Data.Len() > 0 {
		element := eventQueueWrap.Data.Front()
		event := element.Value.(joint.Event)

		if time.Since(event.Timestamp) > 5*time.Minute {
			log.Printf("Deleting event: query=%s, delta=%d, timestamp=%s\n",
				event.Query, event.Delta, event.Timestamp.Format(time.RFC3339))

			queriesMapWrap.Mu.Lock()
			queriesMapWrap.Data[event.Query] -= event.Delta
			queriesMapWrap.Mu.Unlock()

			eventQueueWrap.Data.Remove(element)
		} else {
			break
		}
	}
}

func sortQueries(queriesSortedWrap *joint.QueriesSortedWrapped,
	queriesMapWrap *joint.QueriesMapWrapped) {
	log.Println("Sorting queries...")
	queriesSortedWrap.Mu.Lock()
	defer queriesSortedWrap.Mu.Unlock()

	queriesSortedWrap.Data = queriesSortedWrap.Data[:0]

	queriesMapWrap.Mu.RLock()
	for query, amount := range queriesMapWrap.Data {
		log.Printf("Query: %s, Amount: %d\n", query, amount)
		queriesSortedWrap.Data = append(queriesSortedWrap.Data,
			joint.QueryInfo{Query: query, Amount: amount})
	}
	queriesMapWrap.Mu.RUnlock()

	sort.Slice(queriesSortedWrap.Data, func(i, j int) bool {
		return queriesSortedWrap.Data[i].Amount >
			queriesSortedWrap.Data[j].Amount
	})
}

func LoopEverySecond(
	eventQueueWrap *joint.EventQueueWrapped,
	queriesSortedWrap *joint.QueriesSortedWrapped,
	queriesMapWrap *joint.QueriesMapWrapped,
) {
	secondPassed := time.NewTicker(1 * time.Second)

	processEvents(eventQueueWrap, queriesMapWrap)
	sortQueries(queriesSortedWrap, queriesMapWrap)

	<-secondPassed.C
	LoopEverySecond(eventQueueWrap, queriesSortedWrap, queriesMapWrap)
}
