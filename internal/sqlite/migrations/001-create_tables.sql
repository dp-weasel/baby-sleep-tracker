PRAGMA foreign_keys = ON;

-- Tabla maestra para definir qué podemos trackear
CREATE TABLE IF NOT EXISTS event_types (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,      -- 'SUEÑO_INICIO', 'SUEÑO_FIN', 'COMIDA', etc.
    description TEXT                -- Una breve explicación de qué hace el evento
);

-- Tabla de registros (la que va a crecer)
CREATE TABLE IF NOT EXISTS activity_logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_type_id INTEGER NOT NULL,
    event_time TEXT NOT NULL,       -- ISO8601 (YYYY-MM-DD HH:MM:SS)
    note TEXT,
    created_at TEXT DEFAULT (datetime('now', 'localtime')),
    FOREIGN KEY (event_type_id) REFERENCES event_types(id) ON DELETE RESTRICT
);

-- Índice para que las consultas de historia vuelen
CREATE INDEX IF NOT EXISTS idx_logs_time ON activity_logs(event_time);

