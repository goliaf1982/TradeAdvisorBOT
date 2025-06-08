package database

import (
    "database/sql"
    "fmt"
)

func InitTables() error {
    query := `
    CREATE TABLE IF NOT EXISTS market_data (
	id SERIAL PRIMARY KEY,
	symbol TEXT NOT NULL,
	price NUMERIC NOT NULL,
	timestamp TIMESTAMP DEFAULT NOW()
    );
    `
    _, err := db.Exec(query)
    if err != nil {
	return fmt.Errorf("помилка створення таблиці: %w", err)
    }
    return nil
}
