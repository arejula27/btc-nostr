package main

import (
	"log"
	"strconv"
	"time"

	"github.com/nbd-wtf/go-nostr"
)

func main() {

	nostrConf := NostrConf{
		relays:  []string{"wss://nostr-relay.wlvs.space"},
		privkey: "",
	}
	nostrCtrl := NewNostrController(nostrConf)

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

			content := "Block " + hash + " with height " + strconv.Itoa(block.Height)
			event := *NewNostrEvent(1, content, make(nostr.Tags, 0))
			nostrCtrl.publish(event)

		}

	}

	w := NewWatchdog(interval, cb)
	w.Run()

	time.Sleep(time.Second * 10)

}
