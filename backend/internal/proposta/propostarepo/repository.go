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
	query := `
		INSERT INTO propostas (
			id, postagem_id, remetente_id, destinatario_id, status, excluida
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
		ON CONFLICT (id) DO UPDATE SET
			status = EXCLUDED.status,
			excluida = EXCLUDED.excluida;
	`

	_, err := r.DB.ExecContext(ctx, query,
		proposta.ID,
		proposta.PostagemID,
		proposta.RemetenteID,
		proposta.DestinatarioID,
		proposta.Status,
		proposta.Excluida,
	)

	return err
}

func (r Repository) GetPropostasByID(ctx context.Context, usuarioID string) ([]Proposta, error) {
	query := `
		SELECT id, postagem_id, remetente_id, destinatario_id, status, excluida 
		FROM propostas 
		WHERE remetente_id = $1
	`

	rows, err := r.DB.QueryContext(ctx, query, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var propostas []Proposta
	for rows.Next() {
		var proposta Proposta

		err := rows.Scan(
			&proposta.ID,
			&proposta.PostagemID,
			&proposta.RemetenteID,
			&proposta.DestinatarioID,
			&proposta.Status,
			&proposta.Excluida,
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
