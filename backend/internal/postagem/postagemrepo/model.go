package postagemrepo

import (
	"time"
)

type Post struct {
	ID           string    `json:"id"`
	Titulo       string    `json:"titulo"`
	Descricao    string    `json:"descricao"`
	ImagemBase64 string    `json:"imagem_base64"`
	UserID       string    `json:"user_id"`
	Categoria    string    `json:"categoria"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
