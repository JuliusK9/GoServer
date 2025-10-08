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
