package database

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

var db *sql.DB

func Connect(user, password, dbname string, port int) error {
    psqlInfo := fmt.Sprintf("host=localhost port=%d user=%s password=%s dbname=%s sslmode=disable", port, user, password, dbname)
    var err error
    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
	return err
    }
    return db.Ping()
}

func GetDB() *sql.DB {
    return db
}

func Close() {
    _ = db.Close()
}
