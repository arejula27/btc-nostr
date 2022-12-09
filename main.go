package main

import (
	"log"
	"strconv"
	"time"

	"github.com/nbd-wtf/go-nostr"
)

func main() {
	conf, err := loadConfiguration()
	if err != nil {
		log.Fatal("Can't load configuration, error: ", err)

	}

	//Nostr
	nostrConf := conf.Nostr
	nostrCtrl := NewNostrController(nostrConf)

	//Bitcoin
	bit := NewBitcoinCli(conf.Bitcoin)

	//Watchdog
	interval := time.Second * 15
	//cb is the function which will be call each time the interval ends
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

}
