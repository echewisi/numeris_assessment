package config

import (
    "log"
    "github.com/joho/godotenv"
    "github.com/spf13/viper"
)

// Config structure to hold application configurations
type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    Redis    RedisConfig
	Email    EmailConfig
}

// ServerConfig holds server-related configurations
type ServerConfig struct {
    Port string
}

// DatabaseConfig holds detailed database-related configurations
type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    DSN      string
}

// RedisConfig holds Redis-related configurations
type RedisConfig struct {
    Host     string
    Port     string
    Password string
}

type EmailConfig struct {
	SMTPHost  string
	SMTPPort string
	Username string
	Password string
	From   string
}
// LoadConfig reads configuration from .env, config.yaml, and environment variables
func LoadConfig() *Config {
    // Load .env file first
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using environment variables directly")
    }

    // Configure viper to use config.yaml
    viper.SetConfigFile("configs/config.yaml")
    viper.AutomaticEnv() // Override with environment variables

    if err := viper.ReadInConfig(); err != nil {
        log.Printf("Error reading config file: %v. Falling back to environment variables.", err)
    }

    // Unmarshal configuration into the Config struct
    config := &Config{}
    if err := viper.Unmarshal(config); err != nil {
        log.Fatalf("Unable to decode into struct: %v", err)
    }

    return config
}
