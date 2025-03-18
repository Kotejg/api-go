package rotas

import (
	"api/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

// Rota representa todas as rodas
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// coloca todas as rodas dentro do router
func Configurar(mux *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			mux.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autentica(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			mux.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}

	return mux
}
