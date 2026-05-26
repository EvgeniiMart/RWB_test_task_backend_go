package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/EvgeniiMart/RWB_test_task_backend_go/internal/joint"
)

// Extract and validate N from request URI
func extractN(requestURI string, limit int) (int, string) {
	u, err := url.Parse(requestURI)
	if err != nil {
		return 0, "invalid url"
	}

	q := u.Query()

	nStr := q.Get("N")
	if nStr == "" {
		return 0, "N must be in request"
	}

	n, err := strconv.Atoi(nStr)
	if err != nil {
		return 0, "N should be an integer"
	} else if n <= 0 {
		return 0, "N should be positive"
	} else if n > limit {
		return 0, "N should be less or equal than " + strconv.Itoa(limit)
	}

	return n, ""
}

// Main HTTP handler for incoming topN requests
func RequestHandler(w http.ResponseWriter, r *http.Request,
	queriesSortedWrap *joint.QueriesSortedWrapped, cfg *joint.Config) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	n, errStr := extractN(r.RequestURI, cfg.RequestLimit)
	if errStr != "" {
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	res := assembleAnswer(n, queriesSortedWrap)

	resData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "json error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resData)
	if err != nil {
		panic(err)
	}
}
