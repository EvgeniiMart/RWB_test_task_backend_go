package main

import (
	"container/list"
	"log"
	"net/http"
	"sync"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/consumer"
	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/handler"
	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/looper"
)

func main() {
	// Load config
	var cfg, err = joint.LoadConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Init shared data storages
	var eventQueueWrap = joint.EventQueueWrapped{
		Data: list.New(),
		Mu:   sync.RWMutex{},
	}
	var queriesMapWrap = joint.QueriesMapWrapped{
		Data: make(map[string]int),
		Mu:   sync.RWMutex{},
	}
	var queriesSortedWrap = joint.QueriesSortedWrapped{
		Data: make([]joint.QueryInfo, 0),
		Mu:   sync.RWMutex{},
	}

	// Launch all three subservices
	go looper.LoopEverySecond(&eventQueueWrap, &queriesSortedWrap,
		&queriesMapWrap)

	go consumer.BrokerHandle(&eventQueueWrap, &queriesMapWrap, cfg)

	err = http.ListenAndServe(":8080", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handler.RequestHandler(w, r, &queriesSortedWrap, cfg)
		},
	))
	if err != nil {
		log.Fatal(err)
	}
}
