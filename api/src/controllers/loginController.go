package controllers

import (
	"api/src/authorization"
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// autentica ususario
func Login(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(bodyReq, &usuario); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
	}
	defer db.Close()

	repositorio := repositories.NovoRepositoryDeUsuario(db)
	usuarioSalvoBanco, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
	}
	fmt.Println(usuarioSalvoBanco)

	if err = security.Verify(usuarioSalvoBanco.Senha, usuario.Senha); err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := authorization.CriarToken(usuarioSalvoBanco.ID)

	w.Write([]byte(token))
}
