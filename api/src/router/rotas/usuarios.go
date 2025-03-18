package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Funcao:             controllers.CriarUsuario,
		Metodo:             http.MethodPost,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Funcao:             controllers.BuscarUsuarios,
		Metodo:             http.MethodGet,
		RequerAutenticacao: true,
	}, {
		URI:                "/usuarios/{idUsuario}",
		Funcao:             controllers.BuscarUsuario,
		Metodo:             http.MethodGet,
		RequerAutenticacao: true,
	}, {
		URI:                "/usuarios/{idUsuario}",
		Funcao:             controllers.AlterarUsuario,
		Metodo:             http.MethodPut,
		RequerAutenticacao: true,
	}, {
		URI:                "/usuarios/{idUsuario}",
		Funcao:             controllers.DeletarUsuario,
		Metodo:             http.MethodDelete,
		RequerAutenticacao: true,
	}, {
		URI:                "/usuarios/{idUsuario}/seguir",
		Funcao:             controllers.SeguirUsuario,
		Metodo:             http.MethodPost,
		RequerAutenticacao: true,
	}, {
		URI:                "/usuarios/{idUsuario}/parar-de-seguir",
		Funcao:             controllers.PararDeSeguirUsuario,
		Metodo:             http.MethodPost,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{idUsuario}/parar-de-seguir",
		Funcao:             controllers.BuscarSeguidores,
		Metodo:             http.MethodPost,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{idUsuario}/seguidores",
		Funcao:             controllers.BuscarSeguidores,
		Metodo:             http.MethodGet,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{idUsuario}/seguindo",
		Funcao:             controllers.BuscarSeguindo,
		Metodo:             http.MethodGet,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{idUsuario}/alterar-senha",
		Funcao:             controllers.AlterarSenha,
		Metodo:             http.MethodPost,
		RequerAutenticacao: true,
	},
}
