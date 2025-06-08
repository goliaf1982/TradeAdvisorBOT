package database

import (
    "fmt"
)

func InitTables() error {
    queries := []string{
	`CREATE TABLE IF NOT EXISTS virtual_wallet (
	    id SERIAL PRIMARY KEY,
	    symbol VARCHAR(10) NOT NULL,
	    balance NUMERIC(18,8) NOT NULL DEFAULT 0,
	    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`,
	`CREATE TABLE IF NOT EXISTS virtual_orders (
	    id SERIAL PRIMARY KEY,
	    symbol VARCHAR(10) NOT NULL,
	    side VARCHAR(4) NOT NULL CHECK (side IN ('buy', 'sell')),
	    price NUMERIC(18,8) NOT NULL,
	    quantity NUMERIC(18,8) NOT NULL,
	    commission NUMERIC(18,8) DEFAULT 0,
	    profit NUMERIC(18,8),
	    status VARCHAR(10) NOT NULL DEFAULT 'open',
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	    closed_at TIMESTAMP
	);`,
	`CREATE TABLE IF NOT EXISTS market_data (
	    id SERIAL PRIMARY KEY,
	    symbol VARCHAR(10) NOT NULL,
	    price NUMERIC(18,8) NOT NULL,
	    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`,
    }

    for _, q := range queries {
	_, err := DB.Exec(q)
	if err != nil {
	    return fmt.Errorf("❌ Помилка створення таблиці: %v", err)
	}
    }

    fmt.Println("📦 Таблиці успішно створено або вже існують.")
    return nil
}
