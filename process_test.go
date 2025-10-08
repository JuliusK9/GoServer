package main

import (
	"testing"
)

func TestProcess(t *testing.T) {
	t.Run("create process with name", func(t *testing.T) {
		name := "test-process"
		process := NewProcess(name)

		if process.Name != name {
			t.Errorf("expected name %s, got %s", name, process.Name)
		}

		if process.Running {
			t.Error("new process should not be running")
		}
	})

	t.Run("process start", func(t *testing.T) {
		process := NewProcess("test")

		process.Start()

		if !process.Running {
			t.Error("expected process to be running after start")
		}
	})

	t.Run("running is true before start", func(t *testing.T) {
		process := NewProcess("already-running")
		process.Running = true

		process.Start()

		if !process.Running {
			t.Error("process should remain running")
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
		process1 := NewProcess("process1")
		process2 := NewProcess("process2")
		process3 := NewProcess("process3")

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

	t.Run("recover existent process", func(t *testing.T) {
		store := NewProcessStore()
		originalProcess := NewProcess("retrieve-test")
		originalProcess.Running = true

		store.Add(originalProcess)
		retrieved, exists := store.Get("retrieve-test")
		if !exists {
			t.Error("expected to retrieve existing process")
		}
		if retrieved.Name != "retrieve-test" {
			t.Errorf("expected name 'retrieve-test', got '%s'", retrieved.Name)
		}
		if !retrieved.Running {
			t.Error("expected running state to be preserved")
		}
		if retrieved != originalProcess {
			t.Error("expected to get the same process instance")
		}
	})
}
