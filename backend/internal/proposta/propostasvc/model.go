package propostasvc

import "time"

type PropostasFilter struct {
	UsuarioID string
	Status    *string
	FromTS    *time.Time
	ToTS      *time.Time
}
