package database

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

var db *sql.DB

// Connect — підключається до PostgreSQL
func Connect(user, password, dbname string, port int) error {
    psqlInfo := fmt.Sprintf(
        "host=localhost port=%d user=%s password=%s dbname=%s sslmode=disable",
        port, user, password, dbname,
    )
    var err error
    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        return fmt.Errorf("Connect: %w", err)
    }
    return db.Ping()
}

// InitTables — створює усі необхідні таблиці та індекси
func InitTables() error {
    queries := []string{
        // Ринкові дані
        `CREATE TABLE IF NOT EXISTS market_data (
            id SERIAL PRIMARY KEY,
            symbol TEXT NOT NULL,
            price DOUBLE PRECISION NOT NULL,
            timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
        );`,
        `CREATE INDEX IF NOT EXISTS idx_market_data_symbol_time 
         ON market_data(symbol, timestamp DESC);`,

        // Віртуальні кошти
        `CREATE TABLE IF NOT EXISTS virtual_wallet (
            id SERIAL PRIMARY KEY,
            symbol TEXT NOT NULL,
            balance DOUBLE PRECISION NOT NULL DEFAULT 0,
            updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
        );`,
        `CREATE INDEX IF NOT EXISTS idx_wallet_symbol ON virtual_wallet(symbol);`,

        // Віртуальні ордери
        `CREATE TABLE IF NOT EXISTS virtual_orders (
            id SERIAL PRIMARY KEY,
            symbol TEXT NOT NULL,
            side TEXT CHECK (side IN ('buy', 'sell')) NOT NULL,
            price DOUBLE PRECISION NOT NULL,
            quantity DOUBLE PRECISION NOT NULL,
            commission DOUBLE PRECISION NOT NULL DEFAULT 0,
            profit DOUBLE PRECISION,
            status TEXT NOT NULL DEFAULT 'open',
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            closed_at TIMESTAMP
        );`,
        `CREATE INDEX IF NOT EXISTS idx_orders_symbol_status_time 
         ON virtual_orders(symbol, status, created_at DESC);`,
    }

    for _, q := range queries {
        if _, err := db.Exec(q); err != nil {
            return fmt.Errorf("InitTables error: %w", err)
        }
    }
    return nil
}

// GetDB — повертає з'єднання для інших модулів
func GetDB() *sql.DB {
    return db
}
