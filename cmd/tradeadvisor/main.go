package main

import (
    "fmt"
    "log"
    "time"

    "tradeadvisorbot/internal/binance"
    "tradeadvisorbot/internal/config"
    "tradeadvisorbot/internal/database"
)

func main() {
    cfg := config.LoadConfig()

    if err := database.Connect(cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort); err != nil {
	log.Fatalf("DB connect error: %v", err)
    }
    defer database.Close()

    if err := database.InitTables(); err != nil {
	log.Fatalf("DB init error: %v", err)
    }

    symbols := []string{"btcusdt", "ethusdt"}

    for {
	for _, sym := range symbols {
	    price, err := binance.GetPrice(sym)
	    if err != nil {
		log.Printf("GetPrice error: %v", err)
		continue
	    }
	    err = database.InsertPrice(sym, price)
	    if err != nil {
		log.Printf("InsertPrice error: %v", err)
	    }
	}

	for _, sym := range symbols {
	    profit, err := database.CalculateProfit(sym)
	    if err != nil {
		log.Printf("Profit calc error: %v", err)
		continue
	    }
	    fmt.Printf("[%s] Projected Profit: %.2f USDT\n", sym, profit)
	}

	time.Sleep(30 * time.Second)
    }
}
