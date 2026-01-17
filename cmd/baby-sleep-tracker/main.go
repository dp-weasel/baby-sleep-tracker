package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/dp-weasel/baby-sleep-tracker/internal/application"
	"github.com/dp-weasel/baby-sleep-tracker/internal/infrastructure/sqlite"
)

func main() {
	// Open SQLite database
	db, err := sql.Open("sqlite3", "./baby_sleep.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// Initialize repository
	repo := sqlite.NewEventRepository(db)

	// Wire application services
	registerService := &application.RegisterEventService{
		Store: repo,
	}

	queryService := &application.QueryPeriodsService{
		Reader: repo,
	}

	// For now, we just log that the app is wired correctly
	log.Println("Baby Sleep Tracker backend initialized successfully")

	// Placeholder usage to avoid unused warnings
	_ = registerService
	_ = queryService
}
