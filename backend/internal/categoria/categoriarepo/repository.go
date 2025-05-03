package categoriarepo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		DB: db,
	}
}

func (r Repository) GetCategoryByTitle(ctx context.Context, title string) (Categoria, error) {
	query := `
		SELECT id, titulo
		FROM categorias
		WHERE titulo = $1
	`
	var categoria Categoria

	err := r.DB.QueryRowContext(ctx, query, title).Scan(&categoria.ID, &categoria.Titulo)
	if err != nil {
		if err == sql.ErrNoRows {
			return Categoria{}, fmt.Errorf("categoria com título '%s' não encontrada", title)
		}
		return Categoria{}, fmt.Errorf("erro ao buscar categoria: %w", err)
	}

	return categoria, nil
}

func (r Repository) GetAllCategories(ctx context.Context) ([]Categoria, error) {
	query := `SELECT id, titulo FROM categorias`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []Categoria
	for rows.Next() {
		var categoria Categoria
		if err := rows.Scan(&categoria.ID, &categoria.Titulo); err != nil {
			return nil, err
		}
		categorias = append(categorias, categoria)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categorias, nil
}

func (r Repository) CreateCategories(ctx context.Context, categorias []Categoria) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %w", err)
	}
	defer tx.Rollback()

	checkQuery := `SELECT COUNT(*) FROM categoria WHERE titulo = ?`

	insertQuery := `INSERT INTO categoria (id, titulo) VALUES (?, ?)`

	for _, categoria := range categorias {
		var count int
		err := tx.QueryRowContext(ctx, checkQuery, categoria.Titulo).Scan(&count)
		if err != nil {
			return fmt.Errorf("erro ao verificar categoria: %w", err)
		}

		if count > 0 {
			continue
		}

		id := uuid.New()
		_, err = tx.ExecContext(ctx, insertQuery, id, categoria.Titulo)
		if err != nil {
			return fmt.Errorf("erro ao salvar categoria '%s': %w", categoria.Titulo, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("erro ao confirmar transação: %w", err)
	}

	return nil
}
