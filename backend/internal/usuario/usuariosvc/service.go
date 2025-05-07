package usuariosvc

import (
	"context"
	"escambo/internal/usuario/usuariorepo"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository usuariorepo.Repository
}

func NewService(repository usuariorepo.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) UpsertUsuario(ctx context.Context, usuario usuariorepo.Usuario) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("erro ao gerar hash da senha: %w", err)
	}
	usuario.Senha = string(hashedPassword)

	err = s.repository.UpsertUsuario(ctx, usuario)
	if err != nil {
		return fmt.Errorf("erro ao executar upsert de usuário: %w", err)
	}

	return nil
}
