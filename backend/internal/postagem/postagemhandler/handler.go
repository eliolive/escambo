package postagemhandler

import (
	"context"
	"encoding/json"
	"escambo/internal/postagem/postagemsvc"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type PostagemService interface {
	GetPostagem(ctx context.Context, postID string) (postagemsvc.Postagem, error)
	InsertPostagem(ctx context.Context, post postagemsvc.Postagem) error
}

type Handler struct {
	Service PostagemService
}

func NewHandler(service PostagemService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	post, err := h.Service.GetPostagem(r.Context(), postID)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao buscar postagem: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, fmt.Sprintf("erro ao codificar resposta em JSON: %v", err), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) InsertPostagem(w http.ResponseWriter, r *http.Request) {
	var post postagemsvc.Postagem

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, fmt.Sprintf("erro ao decodificar corpo da requisição: %v", err), http.StatusBadRequest)
		return
	}

	if err := h.Service.InsertPostagem(r.Context(), post); err != nil {
		http.Error(w, fmt.Sprintf("erro ao salvar postagem: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Postagem inserida com sucesso"))
}
