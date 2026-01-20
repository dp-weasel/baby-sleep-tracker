package hypermedia

import (
	"time"

	application "github.com/dp-weasel/baby-sleep-tracker/internal/application/query"
	"github.com/dp-weasel/baby-sleep-tracker/internal/domain"
)

// RootResource represents the hypermedia root document.
type RootResource struct {
	State        domain.State              `json:"state"`
	Since        *time.Time                `json:"since,omitempty"`
	Actions      map[string]Action         `json:"actions,omitempty"`
	RecentCycles []application.RecentCycle `json:"recent_cycles,omitempty"`
}

// Action describes a hypermedia action.
type Action struct {
	Method string      `json:"method"`
	Href   string      `json:"href"`
	Body   interface{} `json:"body"`
}

// RootAssembler builds the root hypermedia representation.
type RootAssembler struct {
	recentCycles *application.RecentCyclesQuery
}

func NewRootAssembler(rc *application.RecentCyclesQuery) *RootAssembler {
	return &RootAssembler{recentCycles: rc}
}

// Assemble builds the root resource based on the last event.
func (a *RootAssembler) Assemble(last *domain.Event) (RootResource, error) {
	resource := RootResource{}

	// Determine state
	switch {
	case last == nil:
		resource.State = domain.StateEmpty
		resource.Actions = map[string]Action{
			"sleep_end": newSleepEndAction(),
		}

	case last.Type == domain.SleepEnd:
		resource.State = domain.StateAwake
		resource.Since = &last.Timestamp
		resource.Actions = map[string]Action{
			"sleep_start": newSleepStartAction(),
		}

	case last.Type == domain.SleepStart:
		resource.State = domain.StateAsleep
		resource.Since = &last.Timestamp
		resource.Actions = map[string]Action{
			"sleep_end": newSleepEndAction(),
		}
	}

	// Attach recent cycles
	cycles, err := a.recentCycles.Last(5)
	if err != nil {
		return RootResource{}, err
	}
	resource.RecentCycles = cycles

	return resource, nil
}

func newSleepStartAction() Action {
	return Action{
		Method: "POST",
		Href:   "/events",
		Body: map[string]string{
			"type":      "sleep_start",
			"timestamp": "{required}",
			"notes":     "{optional}",
		},
	}
}

func newSleepEndAction() Action {
	return Action{
		Method: "POST",
		Href:   "/events",
		Body: map[string]string{
			"type":      "sleep_end",
			"timestamp": "{required}",
			"notes":     "{optional}",
		},
	}
}
