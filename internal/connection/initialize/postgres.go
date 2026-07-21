package initialize

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(databaseURL string) (*gorm.DB, error) { 
	gormDB, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err) 
	}
	return gormDB, nil 
}
