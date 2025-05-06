package postagemsvc

import (
	"context"
	"escambo/internal/postagem/postagemrepo"
	"fmt"

	"github.com/google/uuid"
)

type PostagemRepository interface {
	UpsavePostagem(ctx context.Context, post postagemrepo.Post) error
	GetPostagemByID(ctx context.Context, postID string) (postagemrepo.Post, error)
}

type Service struct {
	PostagemRepo PostagemRepository
}

func NewService(
	postRepo postagemrepo.Repository,
) Service {
	return Service{
		PostagemRepo: postRepo,
	}
}

func (s Service) UpsavePostagem(ctx context.Context, post Postagem) error {
	if err := post.Validate(); err != nil {
		return err
	}

	if post.ID == "" {
		newID, err := uuid.NewRandom()
		if err != nil {
			return fmt.Errorf("erro ao gerar UUID: %w", err)
		}
		post.ID = newID.String()
	}

	err := s.PostagemRepo.UpsavePostagem(ctx, postagemrepo.Post{
		ID:        post.ID,
		Titulo:    post.Titulo,
		Descricao: post.Descricao,
		ImagemURL: post.ImagemURL,
		UserID:    post.UserID,
		Categoria: post.Categoria,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetPostagem(ctx context.Context, postagemID string) (Postagem, error) {
	postagem, err := s.PostagemRepo.GetPostagemByID(ctx, postagemID)
	if err != nil {
		return Postagem{}, err
	}

	return Postagem{
		Titulo:    postagem.Titulo,
		Descricao: postagem.Descricao,
		ImagemURL: postagem.ImagemURL,
		UserID:    postagem.UserID,
		Categoria: postagem.Categoria,
	}, nil
}
