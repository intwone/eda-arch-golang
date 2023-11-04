package gorm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	User    string
	Pass    string
	Host    string
	Port    string
	Name    string
	SSLMode string
}

func NewConn(c *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslmode=%s", c.Host, c.Port, c.Pass, c.User, c.Name, c.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
