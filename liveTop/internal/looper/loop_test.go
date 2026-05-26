package looper

import (
	"container/list"
	"testing"
	"time"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
	"github.com/stretchr/testify/require"
)

func TestProcessEvents(t *testing.T) {
	eventQueue := list.New()

	eventQueue.PushBack(joint.Event{
		Query:     "apple",
		Delta:     5,
		Timestamp: time.Now().Add(-10 * time.Second),
	})
	eventQueue.PushBack(joint.Event{
		Query:     "apple",
		Delta:     50,
		Timestamp: time.Now().Add(-10 * time.Second),
	})
	eventQueue.PushBack(joint.Event{
		Query:     "apple",
		Delta:     500,
		Timestamp: time.Now().Add(100 * time.Second),
	})
	eventQueue.PushBack(joint.Event{
		Query:     "apple",
		Delta:     5000,
		Timestamp: time.Now().Add(-10 * time.Second),
	})
	// Only first two should be applied

	eventQueueWrap := &joint.EventQueueWrapped{
		Data: eventQueue,
	}

	queriesMapWrap := &joint.QueriesMapWrapped{
		Data: map[string]int{
			"apple": 5555,
		},
	}

	cfg, err := joint.LoadConfigFromEnv()
	require.NoError(t, err)
	cfg.ExpirationSeconds = 5

	processEvents(eventQueueWrap, queriesMapWrap, cfg)

	require.Equal(t, 5500, queriesMapWrap.Data["apple"])
	require.Equal(t, 2, eventQueueWrap.Data.Len())
}

func TestSortQueries(t *testing.T) {
	queriesMapWrap := &joint.QueriesMapWrapped{
		Data: map[string]int{
			"banana":     5,
			"apple":      10,
			"orange":     1,
			"melon":      -5,
			"strawberry": 20,
		},
	}

	queriesSortedWrap := &joint.QueriesSortedWrapped{}

	sortQueries(queriesSortedWrap, queriesMapWrap, &joint.Config{})

	require.Equal(t, 5, len(queriesSortedWrap.Data))
	require.Equal(t, "strawberry", queriesSortedWrap.Data[0].Query)
	require.Equal(t, "apple", queriesSortedWrap.Data[1].Query)
	require.Equal(t, "banana", queriesSortedWrap.Data[2].Query)
	require.Equal(t, "orange", queriesSortedWrap.Data[3].Query)
	require.Equal(t, "melon", queriesSortedWrap.Data[4].Query)
}
