package looper

import (
	"sort"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"

	"time"
)

func processEvents(eventQueueWrap *joint.EventQueueWrapped,
	queriesMapWrap *joint.QueriesMapWrapped) {
	eventQueueWrap.Mu.Lock()
	defer eventQueueWrap.Mu.Unlock()

	eventQueue := eventQueueWrap.Data

	for eventQueue.Len() > 0 {
		element := eventQueue.Front()
		event := element.Value.(joint.Event)

		if time.Since(event.Timestamp) > 5*time.Second {
			queriesMapWrap.Mu.Lock()
			queriesMapWrap.Data[event.Query] -= event.Delta
			queriesMapWrap.Mu.Unlock()

			eventQueue.Remove(element)
		} else {
			break
		}
	}
}

func sortQueries(queriesSortedWrap *joint.QueriesSortedWrapped,
	queriesMapWrap *joint.QueriesMapWrapped) {
	queriesSortedWrap.Mu.Lock()
	defer queriesSortedWrap.Mu.Unlock()

	queriesSorted := queriesSortedWrap.Data

	queriesSorted = queriesSorted[:0]

	queriesMapWrap.Mu.RLock()
	for query, amount := range queriesMapWrap.Data {
		queriesSorted = append(queriesSorted,
			joint.QueryInfo{Query: query, Amount: amount})
	}
	queriesMapWrap.Mu.RUnlock()

	sort.Slice(queriesSorted, func(i, j int) bool {
		return queriesSorted[i].Amount > queriesSorted[j].Amount
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
