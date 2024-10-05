package config

import (
	"os"
)

type Config struct{
    
	DatabaseDN string
}

// create function to load configuration
func LoadConfig() *Config{
	
//    return configuration
     return &Config{
		DatabaseDN: getEnv("DATABASE_DN", "homework1.db"),
	 }
}

// function calling env file to config database 
func getEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}