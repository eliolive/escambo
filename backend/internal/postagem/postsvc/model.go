package postsvc

import (
	"errors"
)

type Post struct {
	Titulo          string `json:"titulo"`
	Descricao       string `json:"descricao"`
	ImagemURL       string `json:"imagem_url"`
	UserID          string `json:"user_id"`
	TituloCategoria string `json:"titulo_categoria"`
}

func (p *Post) Validate() error {
	if p.Titulo == "" {
		return errors.New("o título é obrigatório")
	}

	if p.Descricao == "" {
		return errors.New("a descrição é obrigatória")
	}

	if p.ImagemURL == "" {
		return errors.New("a URL da imagem é obrigatória")
	}

	if p.UserID == "" {
		return errors.New("o ID do usuário é obrigatório")
	}

	if p.TituloCategoria == "" {
		return errors.New("o nome da categoria é obrigatório")
	}

	return nil
}
