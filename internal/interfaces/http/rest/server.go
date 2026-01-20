package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dp-weasel/baby-sleep-tracker/internal/application/command"
	"github.com/dp-weasel/baby-sleep-tracker/internal/application/query"
	"github.com/dp-weasel/baby-sleep-tracker/internal/domain"
)

// Server wires HTTP handlers to application services.
type Server struct {
	Register *command.RegisterEventService
	Query    *query.QueryPeriodsService
}

func NewServer(
	register *command.RegisterEventService,
	query *query.QueryPeriodsService,
) *Server {
	return &Server{
		Register: register,
		Query:    query,
	}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/events", s.handleRegisterEvent)
	mux.HandleFunc("/periods", s.handleQueryPeriods)

	return mux
}

// --- Handlers ---

func (s *Server) handleRegisterEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Type      string `json:"type"`
		Timestamp string `json:"timestamp"`
		Notes     string `json:"notas"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	ts, err := time.Parse(time.RFC3339, req.Timestamp)
	if err != nil {
		http.Error(w, "invalid timestamp format", http.StatusBadRequest)
		return
	}

	var eventType domain.EventType
	switch req.Type {
	case "sleep_start":
		eventType = domain.SleepStart
	case "sleep_end":
		eventType = domain.SleepEnd
	default:
		http.Error(w, "invalid event type", http.StatusBadRequest)
		return
	}

	if err := s.Register.Register(eventType, ts, req.Notes); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) handleQueryPeriods(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	periods, err := s.Query.Query(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(periods)
}
