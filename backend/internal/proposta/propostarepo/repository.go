package propostarepo

import (
	"context"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		DB: db,
	}
}

func (r Repository) UpsaveProposta(ctx context.Context, proposta Proposta) error {
	return nil
}
