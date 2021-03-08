package models

import (
	"fmt"
	"time"

	"github.com/clydotron/go-app-test/utils"
)

type EventSource struct {
	Events []EventInfo

	eb *utils.EventBus

	ticker  *time.Ticker
	doneCh  chan bool
	eventId int
}

// NewEventSource
func NewEventSource(eb *utils.EventBus) *EventSource {
	es := &EventSource{eb: eb}
	//hook some additional things up?
	return es
}

// Start ...
func (es *EventSource) Start() {
	//start sending events

	// start a ticker:
	es.ticker = time.NewTicker(2000 * time.Millisecond)
	es.doneCh = make(chan bool)

	go func() {
		for {
			select {
			case <-es.doneCh:
				return
			case <-es.ticker.C:
				es.sendEvent()
			}
		}
	}()
}

// Stop ...
func (es *EventSource) Stop() {
	es.doneCh <- true
}

func (es *EventSource) sendEvent() {

	data := &EventInfo{
		Name:      fmt.Sprintln("Event", es.eventId),
		ID:        fmt.Sprintln(es.eventId),
		TimeStamp: time.Now(),
	}
	es.eb.Publish("event", data)
	es.eventId++
}
