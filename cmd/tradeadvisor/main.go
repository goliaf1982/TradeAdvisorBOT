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
        log.Fatal("⛔ Некоректний порт:", err)
    }

    if err := database.Connect(cfg.DBUser, cfg.DBPassword, cfg.DBName, port); err != nil {
        log.Fatal("❌ Помилка підключення до БД:", err)
    }
    if err := database.InitTables(); err != nil {
        log.Fatal("❌ Помилка ініціалізації таблиць:", err)
    }

    fmt.Println("🚀 TradeAdvisorBOT готовий!")

    symbols := []string{"BTCUSDT", "ETHUSDT"}

    for {
        for _, symbol := range symbols {
            price, err := binance.GetPrice(symbol)
            if err != nil {
                log.Printf("❌ Неможливо отримати ціну %s: %v", symbol, err)
                continue
            }

	    _, err := database.GetDB().Exec(`INSERT INTO prices (symbol, price) VALUES ($1, $2)`, symbol, price)

            if err := database.GetDB().Exec(
                `INSERT INTO market_data(symbol, price) VALUES($1, $2)`, symbol, price,
            ); err != nil {
                log.Printf("❌ Помилка збереження ціни %s: %v", symbol, err)
                continue
            }

            fmt.Printf("💾 %s: %.2f — збережено в БД\n", symbol, price)

            profit, openPrice, found := database.CalculateProfit(symbol, price)
            if found {
                if profit >= 0 {
                    fmt.Printf("📈 Поточний прибуток по %s: %.2f (куплено за %.2f)\n",
                        symbol, profit, openPrice)
                } else {
                    fmt.Printf("📉 Поточний збиток по %s: %.2f (куплено за %.2f)\n",
                        symbol, profit, openPrice)
                }
            }
        }

        fmt.Println("\n📊 Звіт останніх цін:")
        report, err := database.GetLatestPrices()
        if err != nil {
            log.Printf("❌ Помилка отримання звіту: %v", err)
        } else {
            for sym, pr := range report {
                fmt.Printf("🔹 %s: %.2f\n", sym, pr)
            }
        }

        time.Sleep(10 * time.Second)
    }
}
