package propostahandler

import "time"

type GetPropostasParams struct {
	UsuarioID string
	Status    *string
	FromTS    *time.Time
	ToTS      *time.Time
}
