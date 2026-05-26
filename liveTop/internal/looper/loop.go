package looper

import (
	"log"
	"sort"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"

	"time"
)

// Delete all obsolete events from eventQueue and update queriesMap accordingly
func processEvents(eventQueueWrap *joint.EventQueueWrapped,
	queriesMapWrap *joint.QueriesMapWrapped, cfg *joint.Config) {
	eventQueueWrap.Mu.Lock()
	defer eventQueueWrap.Mu.Unlock()

	for eventQueueWrap.Data.Len() > 0 {
		element := eventQueueWrap.Data.Front()
		event := element.Value.(joint.Event)

		if time.Since(event.Timestamp) >
			time.Duration(cfg.ExpirationSeconds)*time.Second {
			queriesMapWrap.Mu.Lock()
			queriesMapWrap.Data[event.Query] -= event.Delta
			queriesMapWrap.Mu.Unlock()

			eventQueueWrap.Data.Remove(element)

			if cfg.Verbose {
				log.Printf("[DEBUG] event expired: query=%s, delta=%d\n",
					event.Query, event.Delta)
			}
		} else {
			break
		}
	}
}

// Rebuild queriesSorted based on queriesMap
func sortQueries(queriesSortedWrap *joint.QueriesSortedWrapped,
	queriesMapWrap *joint.QueriesMapWrapped, cfg *joint.Config) {
	queriesSortedWrap.Mu.Lock()
	defer queriesSortedWrap.Mu.Unlock()

	queriesSortedWrap.Data = queriesSortedWrap.Data[:0]

	queriesMapWrap.Mu.RLock()
	for query, amount := range queriesMapWrap.Data {
		queriesSortedWrap.Data = append(queriesSortedWrap.Data,
			joint.QueryInfo{Query: query, Amount: amount})
	}
	queriesMapWrap.Mu.RUnlock()

	sort.Slice(queriesSortedWrap.Data, func(i, j int) bool {
		return queriesSortedWrap.Data[i].Amount >
			queriesSortedWrap.Data[j].Amount
	})

	if cfg.Verbose {
		log.Printf("[DEBUG] queries sorted: %v\n", queriesSortedWrap.Data)
	}
}

// Every second loop
func LoopEverySecond(
	eventQueueWrap *joint.EventQueueWrapped,
	queriesSortedWrap *joint.QueriesSortedWrapped,
	queriesMapWrap *joint.QueriesMapWrapped,
	cfg *joint.Config,
) {
	secondPassed := time.NewTicker(
		time.Duration(cfg.LoopIntervalSeconds) * time.Second)

	processEvents(eventQueueWrap, queriesMapWrap, cfg)
	sortQueries(queriesSortedWrap, queriesMapWrap, cfg)

	<-secondPassed.C
	LoopEverySecond(eventQueueWrap, queriesSortedWrap, queriesMapWrap, cfg)
}
