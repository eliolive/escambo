package propostahandler

import (
	"context"
	"encoding/json"
	"escambo/internal/proposta/propostarepo"
	"net/http"
)

type PropostaSvc interface {
	GetPropostasByID(ctx context.Context, usuarioID string) ([]propostarepo.Proposta, error)
	UpsaveProposta(ctx context.Context, proposta propostarepo.Proposta) error
}

type Handler struct {
	svc PropostaSvc
}

func NewHandler(svc PropostaSvc) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetPropostasByID(w http.ResponseWriter, r *http.Request) {
	usuarioID := r.URL.Query().Get("usuario_id")
	if usuarioID == "" {
		http.Error(w, "usuario_id é obrigatório", http.StatusBadRequest)
		return
	}

	propostas, err := h.svc.GetPropostasByID(r.Context(), usuarioID)
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
	var proposta propostarepo.Proposta

	if err := json.NewDecoder(r.Body).Decode(&proposta); err != nil {
		http.Error(w, "corpo inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.svc.UpsaveProposta(r.Context(), proposta); err != nil {
		http.Error(w, "erro ao salvar proposta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
