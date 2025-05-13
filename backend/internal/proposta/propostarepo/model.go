package propostarepo

import (
	"errors"
)

type PropostaWriteModel struct {
	PostagemID     string `json:"postagem_id"`
	RemetenteID    string `json:"interessado_id"`
	DestinatarioID string `json:"dono_postagem_id"`
	ImagemBase64   string `json:"imagem_base64"`
	Descricao      string `json:"descricao"`
	Categoria      string `json:"categoria"`
	Nome           string `json:"nome"`
}

type PropostasQueryFilter struct {
	UsuarioID string
	Status    *string
	Tipo      string
}

type Produto struct {
	Nome      string `json:"nome"`
	Categoria string `json:"categoria"`
	Descricao string `json:"descricao"`
	Usuario   string `json:"usuario"`
	Imagem    string `json:"imagem"`
}

type PropostaFormatada struct {
	ProdutoPostagem      Produto `json:"produto_postagem"`
	ProdutoPropostaTroca Produto `json:"produto_proposta_troca"`
	Status               string  `json:"status"`
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
