package main

import (
    "fmt"
    "log"
    "strconv"
    "time"

    "tradeadvisorbot/internal/binance"
    "tradeadvisorbot/internal/config"
    "tradeadvisorbot/internal/database"
)

func main() {
    cfg := config.LoadConfig()

    port, err := strconv.Atoi(cfg.DBPort)
    if err != nil {
	log.Fatal("⛔ Некоректний порт бази даних")
    }

    err = database.Connect(cfg.DBUser, cfg.DBPassword, cfg.DBName, port)
    if err != nil {
	log.Fatal("❌ Помилка підключення:", err)
    }

    err = database.InitTables()
    if err != nil {
	log.Fatal("❌ Помилка ініціалізації таблиць:", err)
    }

    fmt.Println("🚀 TradeAdvisorBOT готовий!")

    symbols := []string{"BTCUSDT", "ETHUSDT"}

    for {
	for _, symbol := range symbols {
	    price, err := binance.GetPrice(symbol)
	    if err != nil {
		log.Printf("❌ Неможливо отримати ціну %s: %v\n", symbol, err)
		continue
	    }
	    err = database.SaveMarketPrice(symbol, price)
	    if err != nil {
		log.Printf("❌ Помилка збереження %s: %v\n", symbol, err)
		continue
	    }
	    fmt.Printf("💾 %s: %.2f USD — збережено в БД\n", symbol, price)

	    // 💹 Прогноз прибутку/збитку на основі останньої купівлі
	    profit, openPrice, found := database.CalculateProfit(symbol, price)
	    if found {
		if profit >= 0 {
		    fmt.Printf("📈 Поточний прибуток по %s: %.2f (Куплено за %.2f)\n", symbol, profit, openPrice)
		} else {
		    fmt.Printf("📉 Поточний збиток по %s: %.2f (Куплено за %.2f)\n", symbol, profit, openPrice)
		}
	    }
	}

	fmt.Println("\n📊 Звіт останніх цін:")
	report, err := database.GetLatestPrices()
	if err != nil {
	    log.Printf("❌ Помилка формування звіту: %v\n", err)
	} else {
	    for symbol, price := range report {
		fmt.Printf("🔹 %s: %.2f USD\n", symbol, price)
	    }
	}

	time.Sleep(10 * time.Second)
    }
}
