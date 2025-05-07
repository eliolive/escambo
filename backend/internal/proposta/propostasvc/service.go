package propostasvc

import (
	"context"
	"escambo/internal/proposta/propostarepo"
)

type PropostaRepository interface {
	UpsaveProposta(ctx context.Context, proposta propostarepo.PropostaWriteModel) error
	GetPropostas(ctx context.Context, filter propostarepo.PropostasQueryFilter) ([]propostarepo.PropostaReadModel, error)
}

type Service struct {
	repo PropostaRepository
}

func NewService(repo PropostaRepository) Service {
	return Service{
		repo: repo,
	}
}

func (s *Service) GetPropostas(ctx context.Context, filter PropostasFilter) ([]propostarepo.PropostaReadModel, error) {
	return s.repo.GetPropostas(ctx, propostarepo.PropostasQueryFilter{
		UsuarioID: filter.UsuarioID,
		Status:    filter.Status,
		FromTS:    filter.FromTS,
		ToTS:      filter.ToTS,
	})
}

func (s *Service) UpsaveProposta(ctx context.Context, proposta propostarepo.PropostaWriteModel) error {
	if err := proposta.Validate(); err != nil {
		return err
	}

	return s.repo.UpsaveProposta(ctx, proposta)
}
