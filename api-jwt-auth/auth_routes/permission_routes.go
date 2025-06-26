package auth_routes

import (
	"net/http"
	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/auth"
	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/auth_controllers"
)

type PermissionRoutes interface {
	Routes() []*Route
}

type permissionRoutesImpl struct {
	permissionsController auth_controllers.PermissionsController
}

func NewPermissionRoutes(permissionsController auth_controllers.PermissionsController) *permissionRoutesImpl {
	return &permissionRoutesImpl{permissionsController}
}

func (r *permissionRoutesImpl) Routes() []*Route {
	return []*Route {
		&Route{
			Path: "/permissions",
			Method: http.MethodPost,
			Handler: auth.IsAuthMiddleware((r.permissionsController.PostPermission)),
		},
		&Route{
			Path: "/permissions/{permission_id}",
			Method: http.MethodGet,
			Handler: auth.IsAuthMiddleware((r.permissionsController.GetPermission)),
		},
		&Route{
			Path: "/permissions",
			Method: http.MethodGet,
			Handler:  auth.IsAuthMiddleware((r.permissionsController.GetPermissions)),
		},
		&Route{
			Path: "/permissions/{permission_id}",
			Method:  http.MethodPut,
			Handler: auth.IsAuthMiddleware((r.permissionsController.PutPermission)),
		},
		&Route{
			Path:   "/permissions/{permission_id}",
			Method:  http.MethodDelete,
			Handler: auth.IsAuthMiddleware((r.permissionsController.DeletePermission)),
		},
	}
}


