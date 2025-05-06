package postagemsvc

import (
	"errors"
)

type Postagem struct {
	ID        string `json:"id"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	ImagemURL string `json:"imagem_url"`
	UserID    string `json:"user_id"`
	Categoria string `json:"categoria"`
}

func (p *Postagem) Validate() error {
	if p.Titulo == "" {
		return errors.New("o título é obrigatório")
	}

	if p.Descricao == "" {
		return errors.New("a descrição é obrigatória")
	}

	if p.ImagemURL == "" {
		return errors.New("a URL da imagem é obrigatória")
	}

	if p.Categoria == "" {
		return errors.New("a categoria é obrigatória")
	}

	return nil
}
