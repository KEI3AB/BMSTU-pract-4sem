package worker

import (
	"net/http"
	"os"
	"strconv"
	"time"
)

var maxConcurrent = 10
var sem = make(chan struct{}, maxConcurrent)
var multiplier time.Duration = 1

func init() {
	if m := os.Getenv("DELAY_MULTIPLIER"); m != "" {
		if val, err := strconv.Atoi(m); err == nil {
			multiplier = time.Duration(val)
		}
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func WorkHandler(w http.ResponseWriter, r *http.Request) {
	sem <- struct{}{}

	defer func() { <-sem }()

	delayStr := r.URL.Query().Get("delay")
	if delay, err := time.ParseDuration(delayStr); err == nil {
		time.Sleep(delay * multiplier)
	}
	w.WriteHeader(http.StatusOK)
}
