# Variables
BINARY_NAME=baby-sleep
DB_NAME=baby_sleep.db
MIGRATIONS_DIR=internal/infrastructure/sqlite/migrations

## help: help: muestra esta ayuda
.PHONY: help
help:
	@echo "Uso:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## db/init: inicializa la base de datos sqlite y corre las migraciones
.PHONY: db/init
db/init:
	@echo "Inicializando base de datos..."
	@sqlite3 $(DB_NAME) < $(MIGRATIONS_DIR)/001-create_tables.sql
	@sqlite3 $(DB_NAME) < $(MIGRATIONS_DIR)/002-seed_event_types.sql
	@echo "Base de datos creada: $(DB_NAME)"

## db/shell: abre la consola de sqlite3
.PHONY: db/shell
db/shell:
	@sqlite3 $(DB_NAME)

## build: compila el binario de Go
.PHONY: build
build:
	@echo "Compilando..."
	@go build -o $(BINARY_NAME) ./cmd/server/main.go

## run: ejecuta la aplicación
.PHONY: run
run: build
	@./$(BINARY_NAME)

## clean: limpia binarios y temporales
.PHONY: clean
clean:
	@rm -f $(BINARY_NAME)
	@echo "Limpieza completada"
