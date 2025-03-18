package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	URI:                "/login",
	Funcao:             controllers.Login,
	Metodo:             http.MethodPost,
	RequerAutenticacao: false,
}
