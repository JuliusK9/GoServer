package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestProcessServer(t *testing.T) {
	store := NewProcessStore()
	server := NewProcessServer(store)

	t.Run("POST /process/ create new process", func(t *testing.T) {
		request := newPostProcessRequest("test-process")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if !store.Exists("test-process") {
			t.Error("expected process 'test-process' to exist in store")
		}
	})

	t.Run("GET /process/list retrun all the processes", func(t *testing.T) {
		process1 := NewProcess("process-1")
		process2 := NewProcess("process-2")

		store.Add(process1)
		store.Add(process2)

		process1.Start()
		request := newGetProcessListRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		body := response.Body.String()
		if !strings.Contains(body, `"name":"process-1"`) {
			t.Errorf("expected body to contain process-1, got %s", body)
		}
		if strings.Contains(body, `"name":"process-2"`) {
			t.Errorf("expected body to contain process-2, got %s", body)
		}
	})

	t.Run("POST /process/ without name", func(t *testing.T) {
		request := newPostProcessRequest("")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("POST /process/ duplicate name", func(t *testing.T) {
		request1 := newPostProcessRequest("duplicate")
		response1 := httptest.NewRecorder()
		server.ServeHTTP(response1, request1)

		assertStatus(t, response1.Code, http.StatusAccepted)

		request2 := newPostProcessRequest("duplicate")
		response2 := httptest.NewRecorder()
		server.ServeHTTP(response2, request2)

		assertStatus(t, response2.Code, http.StatusConflict)
	})

	t.Run("process not found 404", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "///invalid-path///", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestRunServerDoesNotCrashImmediately(t *testing.T) {
	go func() {
		RunServer(":5000")
	}()
	time.Sleep(100 * time.Millisecond)
}

func newPostProcessRequest(name string) *http.Request {
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/process/%s", name), nil)
	return req
}

func newGetProcessListRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/process/list", nil)
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}
}
