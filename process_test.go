package main

import (
	"fmt"
	"testing"
)

func TestProcess(t *testing.T) {
	t.Run("create process with name", func(t *testing.T) {
		name := "test-process"
		process := NewProcess(name, DefaultRunFunc)

		if process.Name != name {
			t.Errorf("expected name %s, got %s", name, process.Name)
		}

		if process.Running {
			t.Error("new process should not be running")
		}
	})

	t.Run("process start", func(t *testing.T) {
		process := NewProcess("test", DefaultRunFunc)

		process.Start()

		if !process.Running {
			t.Error("expected process to be running after start")
		}
	})

	t.Run("running is true before start", func(t *testing.T) {
		process := NewProcess("already-running", DefaultRunFunc)
		process.Running = true

		process.Start()

		if !process.Running {
			t.Error("process should remain running")
		}
	})
	t.Run("process with nil run function", func(t *testing.T) {
		process := &Process{
			"test",
			true,
			nil,
		}
		err := process.Start()

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

		expected := fmt.Sprintf("RunFunc is required for process %s", process.Name)
		if err.Error() != expected {
			t.Errorf("unexpected error: got %q, want %q", err.Error(), expected)
		}
	})
}

func TestStore(t *testing.T) {
	t.Run("empty store", func(t *testing.T) {
		store := NewProcessStore()

		processes := store.List()
		if len(processes) != 0 {
			t.Errorf("expected empty store, got %d processes", len(processes))
		}
	})
	t.Run("add multiple processes", func(t *testing.T) {
		store := NewProcessStore()
		process1 := NewProcess("process1", DefaultRunFunc)
		process2 := NewProcess("process2", DefaultRunFunc)
		process3 := NewProcess("process3", DefaultRunFunc)

		store.Add(process1)
		store.Add(process2)
		store.Add(process3)

		processes := store.List()

		if !store.Exists("process1") {
			t.Error("expected process1 to exist in store after Add")
		}
		if !store.Exists("process2") {
			t.Error("expected process2 to exist in store after Add")
		}
		if !store.Exists("process3") {
			t.Error("expected process3 to exist in store after Add")
		}
		if len(processes) != 3 {
			t.Errorf("expected 3 process in store, got %d", len(processes))
		}
	})
}
