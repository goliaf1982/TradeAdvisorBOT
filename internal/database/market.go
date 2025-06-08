package database

import (
    "fmt"
    "time"
)

type Price struct {
    Symbol    string
    Price     float64
    Timestamp time.Time
}

func GetLatestPrices(symbol string, limit int) ([]Price, error) {
    query := `
	SELECT symbol, price, timestamp
	FROM prices
	WHERE symbol = $1
	ORDER BY timestamp DESC
	LIMIT $2;
    `

    rows, err := db.Query(query, symbol, limit)
    if err != nil {
	return nil, fmt.Errorf("get prices error: %w", err)
    }
    defer rows.Close()

    var prices []Price
    for rows.Next() {
	var p Price
	if err := rows.Scan(&p.Symbol, &p.Price, &p.Timestamp); err != nil {
	    return nil, err
	}
	prices = append(prices, p)
    }
    return prices, nil
}
