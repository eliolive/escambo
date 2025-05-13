package usuariohandler

import (
	"encoding/json"
	"escambo/internal/usuario/usuariorepo"
	"escambo/internal/usuario/usuariosvc"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	usuarioService *usuariosvc.Service
}

func NewHandler(usuarioService *usuariosvc.Service) *Handler {
	return &Handler{usuarioService: usuarioService}
}

func (h *Handler) InsertUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario usuariorepo.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Erro ao decodificar corpo da requisição", http.StatusBadRequest)
		return
	}

	err = h.usuarioService.InsertUsuario(r.Context(), usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário inserido com sucesso"))
}

func (h *Handler) UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var usuario usuariorepo.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Erro ao decodificar corpo da requisição", http.StatusBadRequest)
		return
	}

	err = h.usuarioService.UpdateUsuario(r.Context(), id, usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
