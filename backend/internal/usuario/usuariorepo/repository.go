package usuariorepo

import (
	"context"
	"database/sql"
	"errors"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		DB: db,
	}
}
func (r Repository) InsertUsuario(ctx context.Context, usuario Usuario) error {
	query := `
        INSERT INTO usuarios (nome, email, senha, telefone)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (email) DO NOTHING;
    `
	result, err := r.DB.ExecContext(ctx, query, usuario.Nome, usuario.Email, string(usuario.Senha), usuario.Telefone)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("ja existe um cadastro com esse e-mail")
	}

	return nil
}

func (r Repository) UpdateUsuario(ctx context.Context, id string, usuario Usuario) error {
	query := `
        UPDATE usuarios
        SET nome = $1, telefone = $2, email = $3, updated_at = NOW()
        WHERE id = $4;
    `
	_, err := r.DB.ExecContext(ctx, query, usuario.Nome, usuario.Telefone, usuario.Email, id)
	if err != nil {
		return err
	}
	return nil
}
