package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:                "/publicacoes",
		Funcao:             controllers.CriarPublicacao,
		Metodo:             http.MethodPost,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes",
		Funcao:             controllers.BuscarPublicacoes,
		Metodo:             http.MethodGet,
		RequerAutenticacao: true,
	}, {
		URI:                "/publicacoes/{idPublicacao}",
		Funcao:             controllers.BuscarPublicacao,
		Metodo:             http.MethodGet,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{idPublicacao}",
		Funcao:             controllers.AlterarPublicacao,
		Metodo:             http.MethodPut,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{idPublicacao}",
		Funcao:             controllers.DeletarPublicacao,
		Metodo:             http.MethodDelete,
		RequerAutenticacao: true,
	},
}
