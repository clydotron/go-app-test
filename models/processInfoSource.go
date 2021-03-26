package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/clydotron/go-app-test/utils"
)

type ProcessInfoSource struct {
	eb      *utils.EventBus
	ticker  *time.Ticker
	doneCh  chan bool
	eventId int

	m         sync.RWMutex
	processes map[string]ProcessInfo
}

// NewProcessInfoSource ...
func NewProcessInfoSource(eb *utils.EventBus) *ProcessInfoSource {
	ps := &ProcessInfoSource{
		eb:        eb,
		processes: map[string]ProcessInfo{},
	}
	//hook some additional things up?
	return ps
}

// Start ...
func (ps *ProcessInfoSource) Start() {
	//start sending events

	rand.Seed(time.Now().UnixNano())
	//fmt.Println("ProcessInfoSource >> start")

	// start a ticker:
	ps.ticker = time.NewTicker(1000 * time.Millisecond)
	ps.doneCh = make(chan bool)

	go func() {
		for {
			select {
			case <-ps.doneCh:
				return
			case <-ps.ticker.C:
				ps.SendUpdate()
			}
		}
	}()
}

// AddProcess ...
func (ps *ProcessInfoSource) AddProcess(pi ProcessInfo) {
	ps.m.Lock()
	defer ps.m.Unlock()

	fmt.Println("AddProcess:", pi)
	// check to see if the process is already in the map
	// (we dont really care that much for this example)
	if _, exists := ps.processes[pi.ID]; exists {
		fmt.Println("Process already exists:", pi)
		return
	}
	ps.processes[pi.ID] = pi
}

// RemoveProcess ...
func (ps *ProcessInfoSource) RemoveProcess(id string) {

	ps.m.Lock()
	defer ps.m.Unlock()

	_, ok := ps.processes[id]
	if ok {
		delete(ps.processes, id)
	}
}

// Stop ...
func (ps *ProcessInfoSource) Stop() {
	ps.doneCh <- true
}

func (ps *ProcessInfoSource) SendUpdate() {

	ps.m.Lock()
	defer ps.m.Unlock()

	// for each process in the list, use a random number to determine if cpu usage went up/down/unchangged
	for k, v := range ps.processes {
		delta := 0
		r := rand.Intn(5)
		if r == 1 {
			delta = 1
		} else if r == 2 {
			delta = -1
		}
		c := v.CPU + (delta * 2)
		if c > 100 {
			c = 100
		} else if c < 0 {
			c = 0
		}

		v.CPU = c

		ps.processes[k] = v

		x := &ProcessInfo{
			ID:   v.ID,
			Name: v.Name,
			CPU:  c,
		}

		ps.eb.Publish("PI", x)
	}
}
