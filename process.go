package main

import "time"

type Process struct {
	Name    string
	Running bool
}

func NewProcess(name string) *Process {
	return &Process{
		Name:    name,
		Running: false,
	}
}

func (p *Process) Start() {
	if p.Running {
		return
	}
	p.Running = true

	go func() {
		time.Sleep(1 * time.Minute)
		p.Running = false
	}()
}
