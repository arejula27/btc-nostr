package main

import (
	"fmt"
	"time"

	"github.com/nbd-wtf/go-nostr"
)

// some nostr relay in the wild

type NotrController struct {
	pool *nostr.RelayPool
}

type NostrConf struct {
	relays  []string
	pubkey  string
	privkey string
}

type NostrEvent nostr.Event

func CreateTag(key, value string) *nostr.Tag {

	var tag nostr.Tag = []string{key, value}
	return &tag
}

func NewNostrEvent(kind int, content string, tags nostr.Tags) *NostrEvent {

	return &NostrEvent{
		Kind:      kind,
		Content:   content,
		CreatedAt: time.Now(),
		Tags:      tags,
	}

}

func NewNostrController(conf NostrConf) *NotrController {

	// subscribe to relay
	pool := nostr.NewRelayPool()
	pool.SecretKey = &conf.privkey

	// add a nostr relay to our pool
	for _, relay := range conf.relays {
		errChan := pool.Add(relay, nostr.SimplePolicy{Read: true, Write: true})
		err := <-errChan

		if err != nil {
			fmt.Printf("error calling Add(): %s\n")
		}
	}

	return &NotrController{
		pool: pool,
	}
}

func (ctr *NotrController) publish(e NostrEvent) {

	event, statuses, err := ctr.pool.PublishEvent((*nostr.Event)(&e))
	if err != nil {
		fmt.Printf("error calling NotrController.publish(): %s\n", err.Error())
		return
	}

	StatusProcess(event, statuses)

}

// handle events from out publish events
func StatusProcess(event *nostr.Event, statuses chan nostr.PublishStatus) {
	for status := range statuses {
		switch status.Status {
		case nostr.PublishStatusSent:
			fmt.Printf("Sent event with id %s to '%s'.\n", event.ID, status.Relay)
			return
		case nostr.PublishStatusFailed:
			fmt.Printf("Failed to send event with id %s to '%s'.\n", event.ID, status.Relay)
			return
		case nostr.PublishStatusSucceeded:
			fmt.Printf("Event with id %s seen on '%s'.\n", event.ID, status.Relay)
			return
		}
	}
}
