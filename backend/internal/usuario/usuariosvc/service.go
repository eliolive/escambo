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

func (s *Service) InsertUsuario(ctx context.Context, usuario usuariorepo.Usuario) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("erro ao gerar hash da senha: %w", err)
	}
	usuario.Senha = string(hashedPassword)

	return s.repository.InsertUsuario(ctx, usuario)
}

func (s *Service) UpdateUsuario(ctx context.Context, id string, usuario usuariorepo.Usuario) error {
	return s.repository.UpdateUsuario(ctx, id, usuario)
}
