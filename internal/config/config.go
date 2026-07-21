package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct { 
	Port        string 
	DatabaseURL string 
}

func Load() (Config, error) { 
	port, ok := os.LookupEnv("PORT")
	if !ok || strings.TrimSpace(port) == "" { 
		return Config{}, fmt.Errorf("PORT is required")
	}

	dbURL, ok := os.LookupEnv("DATABASE_URL") 
	if !ok || strings.TrimSpace(dbURL) == "" { 
		return Config{}, fmt.Errorf("DATABASE_URL is required") 
	}

	return Config{ 
		Port:        strings.TrimSpace(port), 
		DatabaseURL: strings.TrimSpace(dbURL), 
	}, nil
}
