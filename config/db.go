package config

import (
    "database/sql"
    "employee/utils"
    "log"
    _"github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
    user := utils.GetEnv("DB_USER", "postgres")
    password := utils.GetEnv("DB_PASSWORD", "123456")
    dbname := utils.GetEnv("DB_NAME", "go_employee")
    dataSourceName := "user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

    var err error
    DB, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatalf("Error connecting to the database: %q", err)
    }

    log.Println("Database connection established")
}

func CloseDB() {
    if DB != nil {
        DB.Close()
    }
}
