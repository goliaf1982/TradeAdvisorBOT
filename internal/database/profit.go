package database

import "fmt"

func CalculateProfit(symbol string) (float64, error) {
    query := `
    SELECT 
	COALESCE(SUM(CASE WHEN side = 'SELL' THEN price * quantity - fee ELSE 0 END), 0) -
	COALESCE(SUM(CASE WHEN side = 'BUY' THEN price * quantity + fee ELSE 0 END), 0)
    FROM orders
    WHERE symbol = $1
    `
    var profit float64
    err := db.QueryRow(query, symbol).Scan(&profit)
    if err != nil {
	return 0, fmt.Errorf("profit query error: %w", err)
    }
    return profit, nil
}
