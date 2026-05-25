package joint

import (
	"container/list"
	"sync"
	"time"
)

type Event struct {
	Query     string    `json:"query"`
	Delta     int       `json:"delta"`
	Timestamp time.Time `json:"timestamp"`
}

type QueryInfo struct {
	Query  string
	Amount int
}

type EventQueueWrapped struct {
	Data *list.List
	Mu   sync.RWMutex
}

type QueriesMapWrapped struct {
	Data map[string]int
	Mu   sync.RWMutex
}

type QueriesSortedWrapped struct {
	Data []QueryInfo
	Mu   sync.RWMutex
}
