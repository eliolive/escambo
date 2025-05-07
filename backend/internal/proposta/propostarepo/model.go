package propostarepo

import (
	"errors"
	"time"
)

type PropostaReadModel struct {
	ID             string    `json:"id"`
	PostagemID     string    `json:"postagem_id"`
	InteressadoID  string    `json:"interessado_id"`
	DonoPostagemID string    `json:"dono_postagem_id"`
	Status         string    `json:"status"`
	ImagemURL      string    `json:"imagem_url"`
	Descricao      string    `json:"descricao"`
	CreatedAt      time.Time `json:"created_at"`
	ExpiresAt      time.Time `json:"expires_at"`
}

type PropostaWriteModel struct {
	PostagemID     string `json:"postagem_id"`
	RemetenteID    string `json:"interessado_id"`
	DestinatarioID string `json:"dono_postagem_id"`
	Status         string `json:"status"`
	ImagemURL      string `json:"imagem_url"`
	Descricao      string `json:"descricao"`
}

type PropostasQueryFilter struct {
	UsuarioID string
	Status    *string
	FromTS    *time.Time
	ToTS      *time.Time
}

func (p PropostaWriteModel) Validate() error {
	if p.RemetenteID == "" {
		return errors.New("remetenteID é obrigatório")
	}
	if p.DestinatarioID == "" {
		return errors.New("destinatarioID é obrigatório")
	}

	return nil
}
