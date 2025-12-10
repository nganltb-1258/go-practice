package utils

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func LoadEnv() {
    err := godotenv.Load();
    if err != nil {
        log.Println("File env not found — using system environment variables")
    }
}

func GetEnv(key, fallback string) string {
    value, env := os.LookupEnv(key);
    if env {
        return value
    }

    return fallback
}
