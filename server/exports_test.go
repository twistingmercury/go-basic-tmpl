package server

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
)

func SetContext(testContext context.Context) {
	ctx = testContext
}

func SetCounter(counter *prometheus.CounterVec) {
	sampleCounter = counter
}
