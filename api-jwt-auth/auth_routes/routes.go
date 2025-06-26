package auth_routes

import (
	"github.com/austineyoyogie/go-hardware-store/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func AuthInstall(router *mux.Router, permissionRoutes PermissionRoutes,
	userRoutes UserRoutes, loginRoutes LoginRoutes) {

	allRoutes := permissionRoutes.Routes()
	allRoutes = append(allRoutes, userRoutes.Routes()...)
	allRoutes = append(allRoutes, loginRoutes.Routes()...)

	for _, route := range allRoutes {
		handler := middlewares.Logger(route.Handler)
		router.HandleFunc(route.Path, handler).Methods(route.Method)
	}
}
