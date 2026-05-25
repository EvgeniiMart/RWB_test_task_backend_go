package livetop

import (
	"log"
	"net/http"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/consumer"
	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/handler"
	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/looper"
)

func main() {
	var eventQueueWrap joint.EventQueueWrapped
	var queriesMapWrap joint.QueriesMapWrapped
	var queriesSortedWrap joint.QueriesSortedWrapped
	limit := 100

	go looper.LoopEverySecond(&eventQueueWrap, &queriesSortedWrap,
		&queriesMapWrap)

	go consumer.BrokerHandle(&eventQueueWrap, &queriesMapWrap)

	err := http.ListenAndServe(":8080", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handler.RequestHandler(w, r, &queriesSortedWrap, limit)
		},
	))
	if err != nil {
		log.Fatal(err)
	}
}
