package main

import "fmt"

type Process struct {
	Name    string
	Running bool
	RunFunc func() error
}

func NewProcess(name string, runFunc func() error) *Process {
	return &Process{
		Name:    name,
		Running: false,
		RunFunc: runFunc,
	}
}

func (p *Process) Start() error {
	if p.RunFunc == nil {
		return fmt.Errorf("RunFunc is required for process %s", p.Name)
	}

	p.Running = true

	go func() {
		p.RunFunc() // Esegue SEMPRE la RunFunc
		p.Running = false
	}()

	return nil
}

type ProcessStore struct {
	processes map[string]*Process
}

func NewProcessStore() *ProcessStore {
	return &ProcessStore{
		processes: map[string]*Process{},
	}
}

func (s *ProcessStore) Add(p *Process) {
	s.processes[p.Name] = p
}

func (s *ProcessStore) List() []*Process {
	//create empty slice with fixed length
	processes := make([]*Process, 0, len(s.processes))
	for _, process := range s.processes {
		processes = append(processes, process)
	}
	return processes
}

func (s *ProcessStore) Exists(name string) bool {
	_, exists := s.processes[name]
	return exists
}
