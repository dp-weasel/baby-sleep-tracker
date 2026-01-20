package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/dp-weasel/baby-sleep-tracker/internal/application/command"
	"github.com/dp-weasel/baby-sleep-tracker/internal/application/query"
	"github.com/dp-weasel/baby-sleep-tracker/internal/infrastructure/sqlite"
	hyper "github.com/dp-weasel/baby-sleep-tracker/internal/interfaces/http/hypermedia"
	rest "github.com/dp-weasel/baby-sleep-tracker/internal/interfaces/http/rest"
)

func main() {
	// --- Database ---
	db, err := sql.Open("sqlite3", "./baby-sleep-app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// --- Infrastructure ---
	eventTypeResolver, err := sqlite.NewEventTypeResolver(db)
	if err != nil {
		log.Fatal(err)
	}

	eventRepo := sqlite.NewEventRepository(db, eventTypeResolver)

	registerService := command.NewRegisterEventService(
		eventRepo,
		eventRepo, // si implementa EventReader también
	)

	queryService := query.NewQueryPeriodsService(eventRepo)

	// --- Application layer ---
	recentCycles := query.NewRecentCyclesQuery(eventRepo)

	// --- HTTP REST (legacy / debug) ---
	restServer := rest.NewServer(registerService, queryService)
	restHandler := restServer.Routes()

	// --- HTTP Hypermedia ---
	rootAssembler := hyper.NewRootAssembler(recentCycles)
	rootHandler := hyper.NewRootHandler(eventRepo, rootAssembler)

	mux := http.NewServeMux()

	// REST endpoints
	mux.Handle("/events", restHandler)
	mux.Handle("/periods", restHandler)

	// Hypermedia root
	mux.Handle("/", rootHandler)

	log.Println("HTTP server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
