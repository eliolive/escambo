package propostahandler

import (
	"context"
	"encoding/json"
	"errors"
	"escambo/internal/proposta/propostarepo"
	"escambo/internal/proposta/propostasvc"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type PropostaSvc interface {
	GetPropostas(ctx context.Context, filter propostasvc.PropostasFilter) ([]propostarepo.PropostaFormatada, error)
	InsertProposta(ctx context.Context, proposta propostarepo.PropostaWriteModel) error
}

type Handler struct {
	svc PropostaSvc
}

func NewHandler(svc PropostaSvc) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetPropostas(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	params, err := getParams(r)
	if err != nil {
		http.Error(w, "parametros invalidos: "+err.Error(), http.StatusBadRequest)
		return
	}

	propostas, err := h.svc.GetPropostas(r.Context(), propostasvc.PropostasFilter{
		UsuarioID: id,
		Status:    params.Status,
		Tipo:      params.Tipo,
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

func (h *Handler) InsertProposta(w http.ResponseWriter, r *http.Request) {
	var proposta propostarepo.PropostaWriteModel

	if err := json.NewDecoder(r.Body).Decode(&proposta); err != nil {
		http.Error(w, "body inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.svc.InsertProposta(r.Context(), proposta); err != nil {
		http.Error(w, "erro ao salvar proposta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getParams(r *http.Request) (GetPropostasParams, error) {
	status := strings.ToLower(r.URL.Query().Get("status"))

	var statusPtr *string
	if status != "" {
		statusPtr = &status
	}

	tipo := strings.ToLower(r.URL.Query().Get("tipo"))
	if tipo == "" {
		return GetPropostasParams{}, errors.New("o parâmetro 'tipo' deve possuir os valores 'enviadas' ou 'recebidas'")
	}

	return GetPropostasParams{
		Status: statusPtr,
		Tipo:   tipo,
	}, nil
}
