package main

import (
	"fmt"
	"time"
)

func main() {

	interval := time.Second * 2
	w := NewWatchdog(interval, func() { fmt.Println("hola") })
	w.Run()

	time.Sleep(time.Second * 10)
}
