package main

import (
	"log"
	"time"
)

func main() {

	interval := time.Second * 30

	rpcCli, err := newRpcClient("localhost", 8332, "bitcoin", "bitcoin", false, interval)
	if err != nil {
		log.Fatal(err)

	}
	cb := func() {

		hash, err := getBestBlockhash(rpcCli)
		if err != nil {
			log.Fatal("Can't obtain best block hash, error: ", err)

		}
		log.Println("Best block hash ", hash)
		block, err := getBlockheader(rpcCli, hash)
		if err != nil {
			log.Fatal("Can't obtain block header, error: ", err)

		}
		log.Println("Block height: ", block.Height)

	}

	w := NewWatchdog(interval, cb)
	w.Run()

	time.Sleep(time.Second * 10)

}
