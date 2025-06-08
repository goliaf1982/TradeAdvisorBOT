package binance

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
)

const baseURL = "https://api.binance.com"

type PriceResponse struct {
    Symbol string `json:"symbol"`
    Price  string `json:"price"`
}

// GetPrice отримує актуальну ціну для пари, наприклад "BTCUSDT"
func GetPrice(symbol string) (float64, error) {
    url := fmt.Sprintf("%s/api/v3/ticker/price?symbol=%s", baseURL, symbol)

    resp, err := http.Get(url)
    if err != nil {
	return 0, fmt.Errorf("помилка запиту: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
	return 0, fmt.Errorf("код відповіді: %d", resp.StatusCode)
    }

    var data PriceResponse
    err = json.NewDecoder(resp.Body).Decode(&data)
    if err != nil {
	return 0, fmt.Errorf("помилка декодування JSON: %w", err)
    }

    price, err := strconv.ParseFloat(data.Price, 64)
    if err != nil {
	return 0, fmt.Errorf("помилка конвертації ціни: %w", err)
    }

    return price, nil
}
