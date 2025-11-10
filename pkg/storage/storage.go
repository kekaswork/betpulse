package storage

import (
	"context"
	"time"
)

type Storage interface {
	Connect(ctx context.Context) error
	Close()
	Ping(ctx context.Context) error
}

type Config struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	DSN      string
	Timeout  time.Duration
}
