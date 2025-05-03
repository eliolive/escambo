package usuariohandler

import (
	"encoding/json"
	"escambo/internal/usuario/usuariorepo"
	"escambo/internal/usuario/usuariosvc"
	"net/http"
)

type Handler struct {
	usuarioService *usuariosvc.Service
}

func NewHandler(usuarioService *usuariosvc.Service) *Handler {
	return &Handler{usuarioService: usuarioService}
}

func (h *Handler) UpsertUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario usuariorepo.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Erro ao decodificar corpo da requisição", http.StatusBadRequest)
		return
	}

	err = h.usuarioService.UpsertUsuario(r.Context(), usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuário inserido/atualizado com sucesso"))
}
