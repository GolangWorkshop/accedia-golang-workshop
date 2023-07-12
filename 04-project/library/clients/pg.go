package clients

import (
	"database/sql"

	"github.com/GolangWorkshop/library/config"
	_ "github.com/lib/pq"
)

func NewPgClient(cfg config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
