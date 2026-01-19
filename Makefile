APP_NAME := baby-sleep-tracker
BIN_DIR := bin
CMD_DIR := ./cmd/$(APP_NAME)
DB_FILE := baby-sleep-app.db
MIGRATIONS_DIR := internal/infrastructure/sqlite/migrations

# Raspberry Pi 3B+ (ARMv7)
PI_GOOS := linux
PI_GOARCH := arm
PI_GOARM := 7

.PHONY: help build run test clean db-reset db-init build-pi

help:
	@echo "Targets disponibles:"
	@echo "  build       Compila el binario para desarrollo local"
	@echo "  run         Compila y ejecuta la app"
	@echo "  test        Ejecuta todos los tests"
	@echo "  clean       Borra binarios"
	@echo "  db-init     Crea la base SQLite y ejecuta migraciones"
	@echo "  db-reset    Borra la base de datos SQLite"
	@echo "  build-pi    Compila el binario para Raspberry Pi 3B+"

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)

run: build
	./$(BIN_DIR)/$(APP_NAME)

test:
	go test ./...

clean:
	rm -rf $(BIN_DIR)

# Inicializa la base de datos local ejecutando migraciones
db-init:
	sqlite3 $(DB_FILE) < $(MIGRATIONS_DIR)/001-create_tables.sql
	sqlite3 $(DB_FILE) < $(MIGRATIONS_DIR)/002-seed_event_types.sql

# ⚠️ Borra la base de datos local. Usar solo en desarrollo.
db-reset:
	rm -f $(DB_FILE)

# Build para Raspberry Pi 3B+ (ARMv7)
build-pi:
	@mkdir -p $(BIN_DIR)
	GOOS=$(PI_GOOS) GOARCH=$(PI_GOARCH) GOARM=$(PI_GOARM) \
		go build -o $(BIN_DIR)/$(APP_NAME)-pi $(CMD_DIR)
