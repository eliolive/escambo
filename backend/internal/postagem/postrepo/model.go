package postrepo

import "time"

type Post struct {
	ID              string    `json:"id"`
	Titulo          string    `json:"titulo"`
	Descricao       string    `json:"descricao"`
	ImagemURL       string    `json:"imagem_url"`
	UserID          string    `json:"user_id"`
	CategoriaID     string    `json:"categoria_id"`
	TituloCategoria string    `json:"categoria_titulo"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
