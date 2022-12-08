package main

import (
	"log"
	"time"
)

func main() {

	interval := time.Second * 15
	timeout := time.Second * 5

	rpcCli, err := newRpcClient("localhost", 8332, "bitcoin", "bitcoin", false, timeout)
	if err != nil {
		log.Fatal(err)

	}

	bit := NewBitcoinCli(rpcCli)
	cb := func() {

		hash, err := bit.getBestBlockhash()
		if err != nil {
			log.Fatal("Can't obtain best block hash, error: ", err)

		}

		if hash != bit.lastBlock {
			bit.lastBlock = hash
			block, err := bit.getBlockheader(hash)
			if err != nil {
				log.Fatal("Can't obtain block header, error: ", err)

			}
			log.Println("Best block hash ", hash)
			log.Println("Block height: ", block.Height)

		}

	}

	w := NewWatchdog(interval, cb)
	w.Run()

	time.Sleep(time.Second * 10)

}
