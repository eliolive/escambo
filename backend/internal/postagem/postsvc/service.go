package postsvc

import (
	"context"
	"escambo/internal/categoria/categoriarepo"
	"escambo/internal/postagem/postrepo"
	"fmt"
)

type PostRepository interface {
	UpsavePost(ctx context.Context, post postrepo.Post) error
	GetPostByID(ctx context.Context, postID string) (postrepo.Post, error)
}

type CategoriaRepository interface {
	GetCategoryByTitle(ctx context.Context, title string) (categoriarepo.Categoria, error)
}

type Service struct {
	PostRepo     PostRepository
	CategoryRepo CategoriaRepository
}

func NewService(
	postRepo postrepo.Repository,
	categoryRepo categoriarepo.Repository,
) Service {
	return Service{
		PostRepo:     postRepo,
		CategoryRepo: categoryRepo,
	}
}

func (s Service) UpsavePost(ctx context.Context, post Post) error {
	if err := post.Validate(); err != nil {
		return err
	}

	categoria, err := s.CategoryRepo.GetCategoryByTitle(ctx, post.TituloCategoria)
	if err != nil {
		return fmt.Errorf("erro ao buscar categoria: %w", err)
	}

	err = s.PostRepo.UpsavePost(ctx, postrepo.Post{
		Titulo:          post.Titulo,
		Descricao:       post.Descricao,
		ImagemURL:       post.ImagemURL,
		UserID:          post.UserID,
		CategoriaID:     categoria.ID,
		TituloCategoria: categoria.Titulo,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetPost(ctx context.Context, postID string) (Post, error) {
	postagem, err := s.PostRepo.GetPostByID(ctx, postID)
	if err != nil {
		return Post{}, err
	}

	return Post{
		Titulo:          postagem.Titulo,
		Descricao:       postagem.Descricao,
		ImagemURL:       postagem.ImagemURL,
		UserID:          postagem.UserID,
		TituloCategoria: postagem.TituloCategoria,
	}, nil
}
