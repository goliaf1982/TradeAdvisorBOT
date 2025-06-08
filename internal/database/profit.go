package database

import (
    "fmt"
)

func CalculateProfit(symbol string) (float64, error) {
    query := `
	SELECT side, price, quantity, fee
	FROM orders
	WHERE symbol = $1
	ORDER BY timestamp ASC;
    `

    rows, err := db.Query(query, symbol)
    if err != nil {
	return 0, fmt.Errorf("profit query error: %w", err)
    }
    defer rows.Close()

    var profit float64
    var totalBuy, totalSell float64

    for rows.Next() {
	var side string
	var price, qty, fee float64

	if err := rows.Scan(&side, &price, &qty, &fee); err != nil {
	    return 0, err
	}

	value := price * qty

	if side == "BUY" {
	    totalBuy += value + fee
	} else if side == "SELL" {
	    totalSell += value - fee
	}
    }

    profit = totalSell - totalBuy
    return profit, nil
}
