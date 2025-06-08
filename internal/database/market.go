package database

import "fmt"

func InsertPrice(symbol string, price float64) error {
    query := `INSERT INTO prices (symbol, price) VALUES ($1, $2)`
    _, err := db.Exec(query, symbol, price)
    if err != nil {
	return fmt.Errorf("insert price error: %w", err)
    }
    return nil
}
