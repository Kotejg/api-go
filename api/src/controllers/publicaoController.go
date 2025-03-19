package controllers

import (
	"api/src/authorization"
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := authorization.ExtrairUsuarioID(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	var publi models.Publicacoes
	if err := json.Unmarshal(body, &publi); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	publi.AutorID = usuarioID

	db, err := banco.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositories.NovoRepositorioPublicacoes(db)
	result, err := repo.Criar(publi)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, result)
}
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}
func AlterarPublicacao(w http.ResponseWriter, r *http.Request) {

}
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
