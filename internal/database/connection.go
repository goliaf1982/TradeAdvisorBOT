package database

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(user, password, dbname string, port int) error {
    psqlInfo := fmt.Sprintf("host=localhost port=%d user=%s password=%s dbname=%s sslmode=disable",
	port, user, password, dbname)

    var err error
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
	return fmt.Errorf("не вдалося відкрити підключення: %v", err)
    }

    err = DB.Ping()
    if err != nil {
	return fmt.Errorf("не вдалося підʼєднатися до БД: %v", err)
    }

    fmt.Println("✅ Підключено до PostgreSQL!")
    return nil
}
