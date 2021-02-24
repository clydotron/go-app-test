package models

import "time"

// TaskInfo ...
type TaskInfo struct {
	Name        string
	Tag         string
	Updated     time.Time
	State       string
	ContainerID string
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
