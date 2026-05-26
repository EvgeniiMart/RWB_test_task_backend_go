package consumer

import (
	"container/list"
	"sync"
	"testing"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
	"github.com/stretchr/testify/require"
)

func TestStoreEvents(t *testing.T) {
	eventQueue := &joint.EventQueueWrapped{
		Data: list.New(),
		Mu:   sync.RWMutex{},
	}

	queriesMap := &joint.QueriesMapWrapped{
		Data: map[string]int{},
		Mu:   sync.RWMutex{},
	}

	events := []joint.Event{
		{
			Query: "apple",
			Delta: 5,
		},
		{
			Query: "apple",
			Delta: 3,
		},
		{
			Query: "banana",
			Delta: -2,
		},
	}

	storeEvents(eventQueue, queriesMap, events)

	require.Equal(t, 8, queriesMap.Data["apple"])
	require.Equal(t, -2, queriesMap.Data["banana"])

	require.Equal(t, 3, eventQueue.Data.Len())
}

func TestJSONValid(t *testing.T) {
	valid := []interface{}{
		map[string]interface{}{
			"query": "apple",
			"delta": 5,
		},
	}

	err := validateJSON("../../data/contract.schema.json", valid)

	require.NoError(t, err)
}

func TestJSONInvalid(t *testing.T) {
	invalid := []interface{}{
		map[string]interface{}{
			"delta": 5,
		},
	}

	err := validateJSON("../../data/contract.schema.json", invalid)

	require.Error(t, err)
	require.Contains(t, err.Error(), "missing properties")
}
