package postagemrepo

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

func (r Repository) UpsavePostagem(ctx context.Context, post Post) error {
	query := `
		INSERT INTO postagens (id, titulo, descricao, imagem_url, user_id, categoria)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id)
		DO UPDATE 
		SET titulo = EXCLUDED.titulo, 
			descricao = EXCLUDED.descricao, 
			imagem_url = EXCLUDED.imagem_url, 
			updated_at = NOW();
	`

	_, err := r.DB.ExecContext(ctx,
		query,
		post.ID,
		post.Titulo,
		post.Descricao,
		post.ImagemURL,
		post.UserID,
		post.Categoria,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetPostagemByID(ctx context.Context, postID string) (Post, error) {
	query := `
		SELECT 
			id, 
			titulo, 
			descricao, 
			imagem_url, 
			user_id, 
			categoria,
			created_at, 
			updated_at
		FROM postagens
		WHERE p.id = $1;
	`
	var post Post
	err := r.DB.QueryRowContext(ctx, query, postID).Scan(
		&post.ID,
		&post.Titulo,
		&post.Descricao,
		&post.ImagemURL,
		&post.UserID,
		&post.Categoria,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (r Repository) GetAllCategories(ctx context.Context) ([]string, error) {
	query := `SELECT DISTINCT categoria FROM postagem`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []string
	for rows.Next() {
		var categoria string
		if err := rows.Scan(&categoria); err != nil {
			return nil, err
		}
		categorias = append(categorias, categoria)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categorias, nil
}
