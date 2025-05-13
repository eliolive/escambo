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

// InsertUsuario godoc
// @Summary      Cadastra um novo usuário
// @Description  Insere um usuário no sistema com base nos dados fornecidos no corpo da requisição
// @Tags         usuarios
// @Accept       json
// @Produce      plain
// @Param        usuario  body  usuariorepo.Usuario  true  "Dados do novo usuário"
// @Success      201  {string}  string  "Usuário inserido com sucesso"
// @Failure      400  {string}  string  "Erro ao decodificar corpo da requisição"
// @Failure      500  {string}  string  "Erro interno do servidor"
// @Router       /usuarios [post]
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

// UpdateUsuario godoc
// @Summary      Atualiza dados de um usuário
// @Description  Atualiza as informações de um usuário identificado pelo ID
// @Tags         usuarios
// @Accept       json
// @Produce      plain
// @Param        id       path   string                true  "ID do usuário"
// @Param        usuario  body   usuariorepo.Usuario   true  "Dados atualizados do usuário"
// @Success      204  {string}  string  "No Content"
// @Failure      400  {string}  string  "Erro ao decodificar corpo da requisição"
// @Failure      500  {string}  string  "Erro interno do servidor"
// @Router       /usuarios/{id} [put]
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
