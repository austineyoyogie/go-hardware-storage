package auth_routes

import (
	"net/http"
	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/auth_controllers"
)

type LoginRoutes interface {
	Routes() []*Route
}

type loginRoutesImpl struct {
	loginController auth_controllers.LoginController
}

func UserLoginRoutes(loginController auth_controllers.LoginController) *loginRoutesImpl {
	return &loginRoutesImpl{loginController}
}

func (r *loginRoutesImpl) Routes() []*Route {
	return []*Route {
		&Route{
			Path: "/login",
			Method: http.MethodPost,
			Handler: r.loginController.PostLogin,
		},
		&Route{
			Path: "/refresh",
			Method: http.MethodPost,
			Handler: r.loginController.RefreshToken,
		},
	}
}