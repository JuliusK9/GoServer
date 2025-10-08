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

func (s *ProcessStore) Get(name string) (*Process, bool) {
	process, exists := s.processes[name]
	return process, exists
}
