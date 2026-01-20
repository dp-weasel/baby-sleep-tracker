package hypermedia

import (
	"encoding/json"
	"net/http"

	"github.com/dp-weasel/baby-sleep-tracker/internal/domain/contracts"
)

// RootHandler handles GET / for the hypermedia API.
type RootHandler struct {
	events    contracts.EventReader
	assembler *RootAssembler
}

func NewRootHandler(events contracts.EventReader, assembler *RootAssembler) *RootHandler {
	return &RootHandler{
		events:    events,
		assembler: assembler,
	}
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	last, err := h.events.Last()
	if err != nil {
		http.Error(w, "failed to load last event", http.StatusInternalServerError)
		return
	}

	resource, err := h.assembler.Assemble(last)
	if err != nil {
		http.Error(w, "failed to assemble root resource", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resource)
}
