package postrepo

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

func (r Repository) UpsavePost(ctx context.Context, post Post) error {
	query := `
		INSERT INTO postagens (id, titulo, descricao, imagem_url, user_id, categoria_id)
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
		post.CategoriaID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetPostByID(ctx context.Context, postID string) (Post, error) {
	query := `
		SELECT 
			p.id, 
			p.titulo, 
			p.descricao, 
			p.imagem_url, 
			p.user_id, 
			p.created_at, 
			p.updated_at,
			c.id AS categoria_id,
			c.titulo AS categoria_titulo
		FROM postagens p
		JOIN categorias c ON p.categoria_id = c.id
		WHERE p.id = $1;
	`
	var post Post
	err := r.DB.QueryRowContext(ctx, query, postID).Scan(
		&post.ID,
		&post.Titulo,
		&post.Descricao,
		&post.ImagemURL,
		&post.UserID,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.CategoriaID,
		&post.TituloCategoria,
	)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}
