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

// GetPropostas godoc
// @Summary      Lista propostas do usuário
// @Description  Retorna propostas enviadas ou recebidas por um usuário com base no tipo e status
// @Tags         trocas
// @Param        id    path     string  true  "ID do usuário"
// @Param        tipo  query    string  true  "Tipo de proposta (enviadas ou recebidas)"
// @Param        status query   string  false "Status da proposta (pendente, aceita, recusada)"
// @Produce      json
// @Success      200  {array}   []propostarepo.PropostaFormatada
// @Failure      500  {string}  string
// @Router       /trocas/{id}/historico [get]
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

// InsertProposta godoc
// @Summary      Cadastra nova proposta
// @Description  Registra uma proposta de troca com base nos dados enviados
// @Tags         trocas
// @Accept       json
// @Produce      json
// @Param        proposta  body   propostarepo.PropostaWriteModel  true  "Dados da proposta"
// @Success      201  {string}  string  "created"
// @Failure      400  {string}  string  "body inválido"
// @Failure      500  {string}  string  "erro ao salvar proposta"
// @Router       /trocas [post]
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
