package main

import (
	"time"
)

// Watchdog is a type whicj let us execute callback funtions
// after a established time interval
type Watchdog struct {
	interval time.Duration
	ticker   *time.Ticker
	callback func()
}

// NewWatchdog creates a wacthdog object
func NewWatchdog(interval time.Duration, callback func()) *Watchdog {

	w := Watchdog{
		interval: interval,
		ticker:   time.NewTicker(interval),
		callback: callback,
	}
	return &w
}

// Stop function prevents de callback function to run another time
func (w *Watchdog) Stop() {
	w.ticker.Stop()
}

// Kick resets the current interval, the wjole interval will pass until the callback
// is run
func (w *Watchdog) Kick() {
	w.ticker.Stop()
	w.ticker.Reset(w.interval)
}

func (w *Watchdog) Run() {

	for _ = range w.ticker.C {
		w.callback()
	}

}
