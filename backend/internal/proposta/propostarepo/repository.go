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

func (r Repository) UpsaveProposta(ctx context.Context, proposta PropostaWriteModel) error {
	query := `
		INSERT INTO propostas (
			postagem_id, interessado_id, dono_postagem_id, status, imagem_url, descricao
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
		ON CONFLICT (id) DO UPDATE SET
			status = EXCLUDED.status;
	`

	_, err := r.DB.ExecContext(ctx, query,
		proposta.PostagemID,
		proposta.RemetenteID,
		proposta.DestinatarioID,
		proposta.Status,
		proposta.ImagemURL,
		proposta.Descricao,
	)

	return err
}

func (r Repository) GetPropostas(ctx context.Context, filter PropostasQueryFilter) ([]PropostaReadModel, error) {
	query := `
		SELECT id, postagem_id, interessado_id, dono_postagem_id, status,
		       imagem_url, descricao, created_at, expires_at
		FROM propostas 
		WHERE dono_postagem_id = $1
	`
	rows, err := r.DB.QueryContext(ctx, query, filter.UsuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var propostas []PropostaReadModel
	for rows.Next() {
		var proposta PropostaReadModel
		err := rows.Scan(
			&proposta.ID,
			&proposta.PostagemID,
			&proposta.InteressadoID,
			&proposta.DonoPostagemID,
			&proposta.Status,
			&proposta.ImagemURL,
			&proposta.Descricao,
			&proposta.CreatedAt,
			&proposta.ExpiresAt,
		)
		if err != nil {
			return nil, err
		}
		propostas = append(propostas, proposta)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return propostas, nil
}
