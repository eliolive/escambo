package propostahandler

import (
	"context"
	"encoding/json"
	"errors"
	"escambo/internal/proposta/propostarepo"
	"escambo/internal/proposta/propostasvc"
	"net/http"
	"time"
)

type PropostaSvc interface {
	GetPropostas(ctx context.Context, filter propostasvc.PropostasFilter) ([]propostarepo.PropostaReadModel, error)
	UpsaveProposta(ctx context.Context, proposta propostarepo.PropostaWriteModel) error
}

type Handler struct {
	svc PropostaSvc
}

func NewHandler(svc PropostaSvc) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetPropostas(w http.ResponseWriter, r *http.Request) {
	params, err := getParams(r)
	if err != nil {
		http.Error(w, "parametros invalidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	propostas, err := h.svc.GetPropostas(r.Context(), propostasvc.PropostasFilter{
		UsuarioID: params.UsuarioID,
		FromTS:    params.FromTS,
		ToTS:      params.ToTS,
		Status:    params.Status,
	})
	if err != nil {
		http.Error(w, "erro ao buscar propostas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(propostas); err != nil {
		http.Error(w, "erro ao codificar resposta: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) UpsaveProposta(w http.ResponseWriter, r *http.Request) {
	var proposta propostarepo.PropostaWriteModel

	if err := json.NewDecoder(r.Body).Decode(&proposta); err != nil {
		http.Error(w, "body inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.svc.UpsaveProposta(r.Context(), proposta); err != nil {
		http.Error(w, "erro ao salvar proposta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getParams(r *http.Request) (GetPropostasParams, error) {
	usuarioID := r.URL.Query().Get("usuario_id")
	if usuarioID == "" {
		return GetPropostasParams{}, errors.New("usuario_id é obrigatório")
	}

	fromTS := r.URL.Query().Get("from_ts")
	var parsedFromTS *time.Time
	if fromTS != "" {
		parsedTime, err := time.Parse(time.RFC3339, fromTS)
		if err != nil {
			return GetPropostasParams{}, errors.New("formato de data 'from_ts' inválido")
		}
		parsedFromTS = &parsedTime
	}

	toTS := r.URL.Query().Get("to_ts")
	var parsedToTS *time.Time
	if toTS != "" {
		parsedTime, err := time.Parse(time.RFC3339, toTS)
		if err != nil {
			return GetPropostasParams{}, errors.New("formato de data 'to_ts' inválido")
		}
		parsedToTS = &parsedTime
	}

	status := r.URL.Query().Get("status")
	if status != "" && status != "pendente" && status != "aceita" && status != "recusada" {
		return GetPropostasParams{}, errors.New("status inválido. Valores válidos: 'pendente', 'aceita', 'recusada'")
	}

	return GetPropostasParams{
		UsuarioID: usuarioID,
		Status:    &status,
		FromTS:    parsedFromTS,
		ToTS:      parsedToTS,
	}, nil
}
