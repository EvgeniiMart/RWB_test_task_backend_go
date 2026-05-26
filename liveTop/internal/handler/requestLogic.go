package handler

import (
	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
)

// Take slice
func getTop(n int, queriesSorted []joint.QueryInfo) []joint.QueryInfo {
	if n > len(queriesSorted) {
		return queriesSorted
	}
	return queriesSorted[:n]
}

// Convert QueryInfo array into string array
func assembleAnswer(n int, queriesSortedWrap *joint.QueriesSortedWrapped) []string {
	queriesSortedWrap.Mu.RLock()
	topK := getTop(n, queriesSortedWrap.Data)
	queriesSortedWrap.Mu.RUnlock()

	result := make([]string, len(topK))
	for i, queryInfo := range topK {
		result[i] = queryInfo.Query
	}

	return result
}
