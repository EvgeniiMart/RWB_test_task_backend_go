package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
	"github.com/stretchr/testify/require"
)

func TestExtractNValid(t *testing.T) {
	n, errStr := extractN("/something?N=5", 10)

	require.Equal(t, "", errStr)
	require.Equal(t, 5, n)
}

func TestExtractNInvalid(t *testing.T) {
	_, errStr := extractN("/something?N=-1", 10)
	require.NotEqual(t, "", errStr)
	require.Contains(t, errStr, "N should be positive")

	_, errStr = extractN("/something?N=abc", 10)
	require.NotEqual(t, "", errStr)
	require.Contains(t, errStr, "N should be an integer")

	_, errStr = extractN("/something?N=2", 1)
	require.NotEqual(t, "", errStr)
	require.Contains(t, errStr, "N should be less or equal than 1")
}

func TestGetTop(t *testing.T) {
	data := []joint.QueryInfo{
		{Query: "a", Amount: 10},
		{Query: "b", Amount: 50},
		{Query: "c", Amount: 100},
		// Notice that amount does not matter at this stage,
		// array is supposed to be already sorted
	}

	top := getTop(2, data)

	require.Equal(t, 2, len(top))
	require.Equal(t, "a", top[0].Query)
	require.Equal(t, "b", top[1].Query)
}

func TestAssembleAnswer(t *testing.T) {
	queriesSortedWrap := &joint.QueriesSortedWrapped{
		Data: []joint.QueryInfo{
			{Query: "apple", Amount: 10},
			{Query: "banana", Amount: 500},
			// Notice that amount does not matter at this stage,
			// array is supposed to be already sorted
		},
	}

	res := assembleAnswer(2, queriesSortedWrap)

	require.Equal(t, []string{"apple", "banana"}, res)
}

func TestRequestHandler(t *testing.T) {
	queriesSortedWrap := &joint.QueriesSortedWrapped{
		Data: []joint.QueryInfo{
			{Query: "apple", Amount: 10},
			{Query: "banana", Amount: 5},
			{Query: "melon", Amount: 1},
		},
	}

	cfg, err := joint.LoadConfigFromEnv()
	// Will be default because no env vars
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/?N=2", nil)
	w := httptest.NewRecorder()

	RequestHandler(w, req, queriesSortedWrap, cfg)

	require.Equal(t, "application/json", w.Header().Get("Content-Type"))
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, `["apple","banana"]`, w.Body.String())
}
