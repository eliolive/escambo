package propostasvc

import (
	"context"
	"escambo/internal/proposta/propostarepo"

	"github.com/google/uuid"
)

type PropostaRepository interface {
	UpsaveProposta(ctx context.Context, proposta propostarepo.Proposta) error
	GetPropostasByID(ctx context.Context, usuarioID string) ([]propostarepo.Proposta, error)
}

type Service struct {
	repo PropostaRepository
}

func NewService(repo PropostaRepository) Service {
	return Service{
		repo: repo,
	}
}

func (s *Service) GetPropostasByID(ctx context.Context, usuarioID string) ([]propostarepo.Proposta, error) {
	return s.repo.GetPropostasByID(ctx, usuarioID)
}

func (s *Service) UpsaveProposta(ctx context.Context, proposta propostarepo.Proposta) error {
	if proposta.ID == "" {
		proposta.ID = uuid.New().String()
	}

	if err := proposta.Validate(); err != nil {
		return err
	}

	return s.repo.UpsaveProposta(ctx, proposta)
}
