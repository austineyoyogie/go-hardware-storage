package product_routes

import (
	"net/http"

	"github.com/austineyoyogie/go-hardware-store/middlewares"
	"github.com/gorilla/mux"		
)

type Route struct {
	Path string
	Method string
	Handler http.HandlerFunc
}

func ProductInstall(router *mux.Router, categoryRoutes CategoryRoutes, productRoutes ProductRoutes) {

	allRoutes := categoryRoutes.Routes()
	allRoutes = append(allRoutes, productRoutes.Routes()...)

	//for _, route := range allRoutes {
	//	router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	//}	
	// it write out every url requested on console.log
	// {"method":"GET","url":"localhost:5000/search/products?q=GTX","version":"HTTP/1.1"}
	// {"method":"GET","url":"localhost:5000/products/1","version":"HTTP/1.1"}

	for _, route := range allRoutes {
		handler := middlewares.Logger(route.Handler)
		router.HandleFunc(route.Path, handler).Methods(route.Method)
	}
}


