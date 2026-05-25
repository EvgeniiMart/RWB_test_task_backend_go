package handler

import "github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"

func getTop(k int, queriesSorted []joint.QueryInfo) []joint.QueryInfo {
	if k > len(queriesSorted) {
		return queriesSorted
	}
	return queriesSorted[:k]
}

func assembleAnswer(k int, queriesSortedWrap *joint.QueriesSortedWrapped) []string {
	queriesSorted := queriesSortedWrap.Data
	queriesSortedWrap.Mu.RLock()
	topK := getTop(k, queriesSorted)
	queriesSortedWrap.Mu.RUnlock()

	result := make([]string, len(topK))
	for i, queryInfo := range topK {
		result[i] = queryInfo.Query
	}

	return result
}
