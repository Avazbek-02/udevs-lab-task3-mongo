package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
    "github.com/spf13/cast"
)

type Config struct {
    MongoDbConnection string
}

func Load() Config {
    if err := godotenv.Load(); err != nil {
        fmt.Printf("Error loading .env file: %v\n", err)
		
    }
    config := Config{
        MongoDbConnection:     cast.ToString(getEnv("MongoDbConnection", "1234")),
    }

    return config
}

func getEnv(key string, defaultVal interface{}) interface{} {
    val, exists := os.LookupEnv(key)
    if !exists {
        return defaultVal
    }
    return val
}
