package propostarepo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		DB: db,
	}
}

func (r Repository) InsertProposta(ctx context.Context, proposta PropostaWriteModel) error {
	checkQuery := `
		SELECT 1 FROM propostas
		WHERE postagem_id = $1 AND interessado_id = $2 AND nome = $3
	`
	var exists int
	err := r.DB.QueryRowContext(ctx, checkQuery,
		proposta.PostagemID,
		proposta.RemetenteID,
		proposta.Nome,
	).Scan(&exists)

	if err == nil {
		return fmt.Errorf("proposta já enviada para essa postagem com esse nome")
	}
	if err != sql.ErrNoRows {
		return err
	}

	query := `
		INSERT INTO propostas (
			postagem_id, interessado_id, dono_postagem_id, imagem_base64, descricao, nome, categoria
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		);
	`

	_, err = r.DB.ExecContext(ctx, query,
		proposta.PostagemID,
		proposta.RemetenteID,
		proposta.DestinatarioID,
		proposta.ImagemBase64,
		proposta.Descricao,
		proposta.Nome,
		proposta.Categoria,
	)

	return err
}

func (r Repository) GetPropostas(ctx context.Context, filter PropostasQueryFilter) ([]PropostaFormatada, error) {
	query := `
		SELECT 
			p.status,

			po.titulo AS postagem_titulo,
			po.descricao AS postagem_descricao,
			po.categoria AS postagem_categoria,
			uo.nome AS postagem_usuario,

			p.nome AS proposta_titulo,
			p.categoria AS proposta_categoria,
			p.descricao AS proposta_descricao,
			ui.nome AS proposta_usuario

		FROM propostas p
		JOIN postagens po ON p.postagem_id = po.id
		JOIN usuarios uo ON po.user_id = uo.id
		JOIN usuarios ui ON p.interessado_id = ui.id
	`

	filteredQuery, args := buildWhere(query, filter)
	rows, err := r.DB.QueryContext(ctx, filteredQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var propostas []PropostaFormatada
	for rows.Next() {
		var p PropostaFormatada
		err := rows.Scan(
			&p.Status,

			&p.ProdutoPostagem.Nome,
			&p.ProdutoPostagem.Descricao,
			&p.ProdutoPostagem.Categoria,
			&p.ProdutoPostagem.Usuario,

			&p.ProdutoPropostaTroca.Nome,
			&p.ProdutoPropostaTroca.Categoria,
			&p.ProdutoPropostaTroca.Descricao,
			&p.ProdutoPropostaTroca.Usuario,
		)
		if err != nil {
			return nil, err
		}
		propostas = append(propostas, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(propostas) == 0 {
		return nil, fmt.Errorf("nenhum dado encontrado")
	}

	return propostas, nil
}

func buildWhere(baseSQL string, filter PropostasQueryFilter) (string, []interface{}) {
	conditions := []string{}
	args := []interface{}{}
	argIndex := 1

	if filter.Tipo == "enviadas" {
		conditions = append(conditions, fmt.Sprintf("interessado_id = $%d", argIndex))
		args = append(args, filter.UsuarioID)
		argIndex++
	} else if filter.Tipo == "recebidas" {
		conditions = append(conditions, fmt.Sprintf("dono_postagem_id = $%d", argIndex))
		args = append(args, filter.UsuarioID)
		argIndex++
	}

	if filter.Status != nil {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argIndex))
		args = append(args, *filter.Status)
		argIndex++
	}

	if len(conditions) > 0 {
		baseSQL += " WHERE " + strings.Join(conditions, " AND ")
	}

	return baseSQL, args
}
