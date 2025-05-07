package usuariorepo

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

func (r Repository) UpsertUsuario(ctx context.Context, usuario Usuario) error {
	query := `
	INSERT INTO usuarios (nome, email, senha, telefone, whatsapp_link)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (id)
	DO UPDATE 
	SET nome = EXCLUDED.nome, 
		email = EXCLUDED.email,
		senha = EXCLUDED.senha, -- Atualizando a senha se já existir
		telefone = EXCLUDED.telefone,
		whatsapp_link = EXCLUDED.whatsapp_link,
		updated_at = NOW();
	`

	_, err := r.DB.ExecContext(ctx, query, usuario.Nome, usuario.Email, string(usuario.Senha), usuario.Telefone, usuario.WhatsappLink)
	if err != nil {
		return err
	}

	return nil
}
