package main

import (
	"time"

	"github.com/trx35479/camunda-autoscaler/autoscaler"
	"github.com/trx35479/camunda-autoscaler/autoscaler/log"
)

var logger = log.NewLogger()

// Main handler of the function
func handler(t time.Time) {
	if err := autoscaler.Handler(); err != nil {
		logger.Fatal("autoscaler returns error, shutting down.")
	}
}

// Polling function to initiate a polling to api
func polling(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func main() {
	polling(10*time.Second, handler)
}
