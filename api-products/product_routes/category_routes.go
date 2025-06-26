package product_routes

import (
	"net/http"

	"github.com/austineyoyogie/go-hardware-store/api-products/product_controllers"
	"github.com/austineyoyogie/go-hardware-store/middlewares"
)

type CategoryRoutes interface {
	Routes() []*Route
}

type categoryRoutesImpl struct {
	categoriesController product_controllers.CategoriesController
}

func NewCategoryRoutes(categoriesController product_controllers.CategoriesController) *categoryRoutesImpl {
	return &categoryRoutesImpl{categoriesController}
}

func (r *categoryRoutesImpl) Routes() []*Route {
	return []*Route {
		&Route{
			Path: "/categories",
			Method: http.MethodPost,
			Handler: r.categoriesController.PostCategory,
		},
		&Route{
			Path: "/categories",
			Method: http.MethodGet,
			// I can also use middleware like this to restrict some access
			Handler: middlewares.Logger(r.categoriesController.GetCategories),
		},
		&Route{
			Path: "/categories/{category_id}",
			Method: http.MethodGet,
			Handler: r.categoriesController.GetCategory,
		},
		&Route{
			Path:    "/categories/{category_id}",
			Method:  http.MethodPut,
			Handler: r.categoriesController.PutCategory,
		},
		&Route{
			Path:    "/categories/{category_id}",
			Method:  http.MethodDelete,
			Handler: r.categoriesController.DeleteCategory,
		},
	}
}

