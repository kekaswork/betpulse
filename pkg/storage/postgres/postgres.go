package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kekaswork/betpulse/pkg/storage"
)

type Postgres struct {
	cfg  storage.Config
	pool *pgxpool.Pool
}

func NewPostgres(cfg storage.Config) (*Postgres, error) {
	return &Postgres{cfg: cfg}, nil
}

func (p *Postgres) Connect(ctx context.Context) error {
	dsn := p.cfg.DSN
	if dsn == "" {
		dsn = fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
			p.cfg.User,
			p.cfg.Password,
			p.cfg.Host,
			p.cfg.Port,
			p.cfg.DbName,
		)
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return fmt.Errorf("connect to postgres: %w", err)
	}

	ctxPing, cancel := context.WithTimeout(ctx, p.cfg.Timeout)
	defer cancel()

	if err := pool.Ping(ctxPing); err != nil {
		pool.Close()
		return fmt.Errorf("ping: %w", err)
	}

	p.pool = pool
	fmt.Println("Connected to Postgres successfully!")
	return nil
}

func (p *Postgres) Close() {
	if p.pool != nil {
		p.pool.Close()
		fmt.Println("Postgres connection closed.")
	}
}

func (p *Postgres) Ping(ctx context.Context) error {
	return p.Ping(ctx)
}
