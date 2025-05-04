package propostarepo

import "errors"

type Proposta struct {
	ID             string `json:"id"`
	PostagemID     string `json:"postagem_id"`
	RemetenteID    string `json:"remetente_id"`
	DestinatarioID string `json:"destinatario_id"`
	Status         string `json:"status"`
	Excluida       bool   `json:"excluida"`
}

func (p Proposta) Validate() error {
	if p.RemetenteID == "" {
		return errors.New("remetenteID é obrigatório")
	}
	if p.DestinatarioID == "" {
		return errors.New("destinatarioID é obrigatório")
	}

	return nil
}
