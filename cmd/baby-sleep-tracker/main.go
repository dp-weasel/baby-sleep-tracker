package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/dp-weasel/baby-sleep-tracker/internal/application"
	"github.com/dp-weasel/baby-sleep-tracker/internal/infrastructure/sqlite"
	httpapi "github.com/dp-weasel/baby-sleep-tracker/internal/interfaces/http"
)

func main() {
	// Open SQLite database
	db, err := sql.Open("sqlite3", "./baby-sleep-app.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// Initialize EventTypeResolver (fail fast on misconfigured DB)
	resolver, err := sqlite.NewEventTypeResolver(db)
	if err != nil {
		log.Fatalf("failed to initialize event types resolver: %v", err)
	}

	// Initialize repository
	repo := sqlite.NewEventRepository(db, resolver)

	// Wire application services
	registerService := &application.RegisterEventService{
		Store: repo,
	}

	queryService := &application.QueryPeriodsService{
		Reader: repo,
	}

	// For now, we just log that the app is wired correctly
	log.Println("Baby Sleep Tracker backend initialized successfully")

	// Initialize HTTP server
	server := httpapi.NewServer(registerService, queryService)

	addr := ":8080"
	log.Printf("HTTP server listening on %s", addr)

	if err := http.ListenAndServe(addr, server.Routes()); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
