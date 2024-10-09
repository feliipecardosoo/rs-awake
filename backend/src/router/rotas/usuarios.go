package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuario = []Rota{
	{
		URI:                  "/usuario",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CriarUsuario,
		RequerAutentificacao: false,
	},
	{
		URI:                  "/usuarios",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarUsuarios,
		RequerAutentificacao: false,
	},
	{
		URI:                  "/usuarios/{usuarioId}",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarUsuario,
		RequerAutentificacao: false,
	},
	{
		URI:                  "/usuarios/{usuarioId}",
		Metodo:               http.MethodPut,
		Funcao:               controllers.AtualizarUsuario,
		RequerAutentificacao: false,
	},
	{
		URI:                  "/usuarios/{usuarioId}",
		Metodo:               http.MethodDelete,
		Funcao:               controllers.DeletarUsuario,
		RequerAutentificacao: false,
	},
}
