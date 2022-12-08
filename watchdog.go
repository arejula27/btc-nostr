package main

import (
	"time"
)

// Watchdog is a type whicj let us execute callback funtions
// after a established time interval
type watchdog struct {
	interval time.Duration
	ticker   *time.Ticker
	callback func()
}

// NewWatchdog creates a wacthdog object
func NewWatchdog(interval time.Duration, callback func()) *watchdog {

	w := watchdog{
		interval: interval,
		ticker:   time.NewTicker(interval),
		callback: callback,
	}
	return &w
}

// Stop function prevents de callback function to run another time
func (w *watchdog) Stop() {
	w.ticker.Stop()
}

// Kick resets the current interval, the wjole interval will pass until the callback
// is run
func (w *watchdog) Kick() {
	w.ticker.Stop()
	w.ticker.Reset(w.interval)
}

func (w *watchdog) Run() {

	for range w.ticker.C {
		w.callback()
	}

}
