package database

import (
    "fmt"
)

func InitTables() error {
    tableQueries := []string{
	`CREATE TABLE IF NOT EXISTS prices (
	    id SERIAL PRIMARY KEY,
	    symbol TEXT NOT NULL,
	    price DOUBLE PRECISION NOT NULL,
	    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`,

	`CREATE INDEX IF NOT EXISTS idx_prices_symbol_time ON prices(symbol, timestamp DESC);`,

	`CREATE TABLE IF NOT EXISTS orders (
	    id SERIAL PRIMARY KEY,
	    symbol TEXT NOT NULL,
	    side TEXT CHECK (side IN ('BUY', 'SELL')) NOT NULL,
	    price DOUBLE PRECISION NOT NULL,
	    quantity DOUBLE PRECISION NOT NULL,
	    fee DOUBLE PRECISION NOT NULL,
	    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`,

	`CREATE INDEX IF NOT EXISTS idx_orders_symbol_time ON orders(symbol, timestamp DESC);`,
    }

    for _, query := range tableQueries {
	_, err := db.Exec(query)
	if err != nil {
	    return fmt.Errorf("DB Init error: %w", err)
	}
    }

    return nil
}
