package propostasvc

import (
	"context"
	"escambo/internal/proposta/propostarepo"
)

type PropostaRepository interface {
	InsertProposta(ctx context.Context, proposta propostarepo.PropostaWriteModel) error
	GetPropostas(ctx context.Context, filter propostarepo.PropostasQueryFilter) ([]propostarepo.PropostaFormatada, error)
}

type Service struct {
	repo PropostaRepository
}

func NewService(repo PropostaRepository) Service {
	return Service{
		repo: repo,
	}
}

func (s *Service) GetPropostas(ctx context.Context, filter PropostasFilter) ([]propostarepo.PropostaFormatada, error) {
	return s.repo.GetPropostas(ctx, propostarepo.PropostasQueryFilter{
		UsuarioID: filter.UsuarioID,
		Status:    filter.Status, //opcional
		Tipo:      filter.Tipo,   //obrigatorio
	})
}

func (s *Service) InsertProposta(ctx context.Context, proposta propostarepo.PropostaWriteModel) error {
	if err := proposta.Validate(); err != nil {
		return err
	}

	return s.repo.InsertProposta(ctx, proposta)
}
