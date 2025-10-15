package process

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type ProcessServer struct {
	store *ProcessStore
}

func NewProcessServer(store *ProcessStore) *ProcessServer {
	return &ProcessServer{store: store}
}

func (s *ProcessServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/process/"):
		s.createProcess(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/process/list"):
		s.listProcesses(w, r)
	default:
		http.NotFound(w, r)
	}
}

func DefaultRunFunc() error {
	time.Sleep(1 * time.Minute)
	return nil
}

func (s *ProcessServer) createProcess(w http.ResponseWriter, r *http.Request) {
	//estrae parametro "name"
	name := strings.TrimPrefix(r.URL.Path, "/process/")
	if name == "" {
		http.Error(w, "name parameter is required", http.StatusBadRequest)
		return
	}
	if s.store.Exists(name) {
		http.Error(w, "process with this name alredy exists", http.StatusConflict)
		return
	}

	process := NewProcess(name, DefaultRunFunc)
	s.store.Add(process)
	process.Start()

	w.WriteHeader(http.StatusAccepted)
}

func (s *ProcessServer) listProcesses(w http.ResponseWriter, r *http.Request) {
	processes := s.store.List()
	var response []string
	for _, p := range processes {
		if p.Running {
			response = append(response, p.Name)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
