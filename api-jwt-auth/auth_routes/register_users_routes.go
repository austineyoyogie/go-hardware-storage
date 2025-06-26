package auth_routes

import (
	"net/http"
	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/auth_controllers"
)

type UserRoutes interface {
	Routes() []*Route
}

type userRoutesImpl struct {
	usersController auth_controllers.UsersController
}

func NewUserRoutes(usersController auth_controllers.UsersController) *userRoutesImpl {
	return &userRoutesImpl{usersController}
}

func (r *userRoutesImpl) Routes() []*Route {
	return []*Route {
		&Route{
			Path: "/users",
			Method: http.MethodPost,
			Handler: r.usersController.PostUser,
		},
		&Route{
			Path: "/verify",
			Method: http.MethodGet,
			Handler: r.usersController.VerifyUser,
		},
		&Route{
			Path:    "/users/{user_id}",
			Method:  http.MethodGet,
			Handler: r.usersController.GetUser,
		},
		&Route{
			Path:    "/users",
			Method:  http.MethodGet,
			Handler: r.usersController.GetUsers,
		},
		&Route{
			Path: "/users/{user_id}",
			Method: http.MethodPut,
			Handler: r.usersController.PutUser,
		},
		&Route{
			Path:   "/users/{user_id}",
			Method:  http.MethodDelete,
			Handler: r.usersController.DeleteUser,
		},
		&Route{
			Path:   "/reset", // PUT localhost:8000/reset
			Method:  http.MethodPut,
			Handler: r.usersController.ResetPasswordUser,
		},
		&Route{
			Path:   "/resetpassword",
			Method:  http.MethodGet,
			Handler: r.usersController.GetNewPasswordUser,
		},
		 &Route{
		 	Path:   "/putpassword", // PUT localhost:8000/putpassword
		 	Method:  http.MethodPut,
		 	Handler: r.usersController.PutNewPasswordUser,
		},
	}
}
