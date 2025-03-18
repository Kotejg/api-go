package controllers

import (
	"api/src/authorization"
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
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
	var userID int
	userID, err := strconv.Atoi(params["idUsuario"])

	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	usuarioIDToken, err := authorization.ExtrairUsuarioID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
	}

	if usuarioIDToken != uint64(userID) {
		responses.Erro(w, http.StatusForbidden,
			errors.New("ação nao autorizada. não é possivel deletar um usuario que não seja o seu"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
	}
	defer db.Close()
	repositorio := repositories.NovoRepositoryDeUsuario(db)
	if erro := repositorio.Deletar(userID); erro != nil {
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
	usuarioIDToken, err := authorization.ExtrairUsuarioID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
	}

	if usuarioIDToken != uint64(userId) {
		responses.Erro(w, http.StatusForbidden, errors.New("ação nao autorizada. não é possivel atualizar um ususario que não seja o seus"))
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
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)

}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	usuarioIDToken, err := authorization.ExtrairUsuarioID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	usuarioID, err := strconv.Atoi(params["idUsuario"])
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if usuarioIDToken == uint64(usuarioID) {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possivel seguir a si mesmo "))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repositorio := repositories.NovoRepositoryDeUsuario(db)
	if err := repositorio.Seguir(usuarioIDToken, uint64(usuarioID)); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// para de seguir outro ususario
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	usuarioIDToken, err := authorization.ExtrairUsuarioID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	params := mux.Vars(r)
	usuarioID, err := strconv.Atoi(params["idUsuario"])
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	if usuarioIDToken == uint64(usuarioID) {
		responses.Erro(w, http.StatusForbidden, errors.New("vc não pode para de seguir a si mesmo"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repositorio := repositories.NovoRepositoryDeUsuario(db)
	if err := repositorio.PararDeSeguir(usuarioIDToken, uint64(usuarioID)); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioID, err := strconv.Atoi(params["idUsuario"])
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repositorio := repositories.NovoRepositoryDeUsuario(db)
	seguidores, err := repositorio.BuscarSeguidores(usuarioID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, seguidores)

}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioID, err := strconv.Atoi(params["idUsuario"])
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repositorio := repositories.NovoRepositoryDeUsuario(db)
	seguidores, err := repositorio.BuscarSeguindo(usuarioID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, seguidores)
}

// altera senha do usuario
func AlterarSenha(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioID, err := strconv.Atoi(params["idUsuario"])
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	usuarioIDToken, err := authorization.ExtrairUsuarioID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioIDToken != uint64(usuarioID) {
		responses.Erro(w, http.StatusForbidden,
			errors.New("não é possivel atualizar um usuario que não seja seu"))
		return
	}

	bodyReq, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
	}
	var senha models.Senha
	if err := json.Unmarshal(bodyReq, &senha); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositoryDeUsuario(db)
	senhaSalvaNoBanco, err := repositorio.BuscarSenhaPorID(usuarioID)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Verify(senhaSalvaNoBanco, senha.OldPassword); err != nil {
		responses.Erro(w, http.StatusUnauthorized, errors.New("Senha nao condiz com que esta salva no banco"))
		return
	}
	senhaComHash, erro := security.Hash(senha.NewPassword)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
