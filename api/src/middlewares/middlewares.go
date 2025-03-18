package middlewares

import (
	"api/src/authorization"
	"api/src/responses"
	"log"
	"net/http"
)

func Autenticar(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := authorization.ValidarToken(r); erro != nil {
			responses.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		nextFunc(w, r)
	}
}

// logger escreve as informacoes da req no terminal
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}
