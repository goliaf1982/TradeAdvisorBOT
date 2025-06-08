package database

import "fmt"

func InsertOrder(symbol, side string, price, quantity, fee float64) error {
    query := `INSERT INTO orders (symbol, side, price, quantity, fee) VALUES ($1, $2, $3, $4, $5)`
    _, err := db.Exec(query, symbol, side, price, quantity, fee)
    if err != nil {
	return fmt.Errorf("insert order error: %w", err)
    }
    return nil
}
