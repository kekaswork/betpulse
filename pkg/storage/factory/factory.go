package factory

import (
	"fmt"

	"github.com/kekaswork/betpulse/pkg/storage"
	"github.com/kekaswork/betpulse/pkg/storage/postgres"
)

func New(cfg storage.Config) (storage.Storage, error) {
	switch cfg.Driver {
	case "postgres", "postgresql":
		return postgres.NewPostgres(cfg)
	// case "mysql":
	// 	return newMySQL(cfg)
	default:
		return nil, fmt.Errorf("unsupported storage driver: %s", cfg.Driver)
	}
}
