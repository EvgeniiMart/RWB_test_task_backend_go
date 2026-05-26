package handler

import (
	"log"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
)

func getTop(n int, queriesSorted []joint.QueryInfo) []joint.QueryInfo {
	if n > len(queriesSorted) {
		return queriesSorted
	}
	return queriesSorted[:n]
}

func assembleAnswer(n int, queriesSortedWrap *joint.QueriesSortedWrapped) []string {
	log.Printf("Len: %d", len(queriesSortedWrap.Data))

	queriesSortedWrap.Mu.RLock()
	topK := getTop(n, queriesSortedWrap.Data)
	queriesSortedWrap.Mu.RUnlock()

	result := make([]string, len(topK))
	for i, queryInfo := range topK {
		result[i] = queryInfo.Query
	}

	return result
}
