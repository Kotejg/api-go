package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// cria ususario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var usuario models.Usuario
	if err = json.Unmarshal(bodyReq, &usuario); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	etapa := "cadastro"

	if err = usuario.PrepararUsuario(etapa); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositoryDeUsuario(db)
	usuario.ID, err = repositorio.Criar(usuario)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, usuario)

}

// busca usuario via ID
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["idUsuario"])
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer db.Close()
	repositorio := repositories.NovoRepositoryDeUsuario(db)
	usuario, err := repositorio.BuscarPorId(uint(userId))
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, usuario)
}

// busca todos os usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	filter := strings.ToLower(r.URL.Query().Get("filter"))
	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repositorio := repositories.NovoRepositoryDeUsuario(db)
	listaUsuarioFiltrados, erro := repositorio.Buscar(filter)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, listaUsuarioFiltrados)
}

// deleta ususario
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ID int
	ID, err := strconv.Atoi(params["idUsuario"])

	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
	}
	defer db.Close()
	repositorio := repositories.NovoRepositoryDeUsuario(db)
	if err := repositorio.Deletar(ID); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)

}

// alterar usuario
func AlterarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["idUsuario"])
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	bodyReq, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(bodyReq, &usuario); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	etapa := "update"
	if erro := usuario.PrepararUsuario(etapa); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repositorio := repositories.NovoRepositoryDeUsuario(db)

	err = repositorio.Atualizar(userId, usuario)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusNoContent, nil)

}
