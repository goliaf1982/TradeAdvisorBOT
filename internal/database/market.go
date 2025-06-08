package database

import "fmt"

// SaveMarketPrice зберігає ціну для символу
func SaveMarketPrice(symbol string, price float64) error {
    query := `
	INSERT INTO market_data (symbol, price)
	VALUES ($1, $2);
    `
    _, err := db.Exec(query, symbol, price)
    if err != nil {
	return fmt.Errorf("помилка збереження ринкової ціни: %w", err)
    }
    return nil
}
