package models

import (
	"time"
)

type Observer interface {
	Updated(i interface{})
	GetID() string
}

// TaskInfo ...
type TaskInfo struct {
	Name        string
	Tag         string
	Updated     time.Time
	State       string
	ContainerID string

	observers []Observer //trying out two different ways, pick one (eventually)
	uc        chan<- bool
}

// SetChannel
func (ti *TaskInfo) SetChannel(ch chan<- bool) {
	ti.uc = ch
}

func (ti *TaskInfo) AddObserver(o Observer) {
	ti.observers = append(ti.observers, o)
}

func (ti *TaskInfo) RemoveObserver(o Observer) {

	for i, oo := range ti.observers {
		if o.GetID() == oo.GetID() {
			ti.observers = append(ti.observers[:i], ti.observers[i+1:]...)
		}
	}
}

func (ti *TaskInfo) NotifyAll() {
	//fmt.Println("NotifyAll")
	for _, o := range ti.observers {
		o.Updated(ti)
	}

	if ti.uc != nil {
		ti.uc <- true
	}
}

//@todo how to do an enum in

// MachineInfo ...
type MachineInfo struct {
	Name   string
	Role   string
	Status string
	Memory int32
	Tasks  []TaskInfo
}

// Events ...
type EventInfo struct {
	Name      string
	ID        string
	Paylod    interface{}
	TimeStamp time.Time
}

type EventInfoRequest struct {
	StartTime time.Time
	Callback  func(events []EventInfo)
}

type ControlPlaneInfo struct {
	Name   string
	Status string
}

type WorkerNodeInfo struct {
	Name   string
	Status string
}

type ClusterInfo struct {
	ControlPlanes map[string]ControlPlaneInfo
	WorkerNodes   map[string]WorkerNodeInfo
}

type ProcessInfo struct {
	Name string
	ID   string
	CPU  int
}
