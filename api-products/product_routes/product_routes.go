package product_routes

import (
	"net/http"
	"github.com/austineyoyogie/go-hardware-store/api-products/product_controllers"
)

type ProductRoutes interface {
	Routes() []*Route
}

type productRoutesImpl struct {
	productsController product_controllers.ProductsController
}

func NewProductRoutes(productsController product_controllers.ProductsController) *productRoutesImpl {
	return &productRoutesImpl{productsController}
}

func (r *productRoutesImpl) Routes() []*Route {
	return []*Route {
		&Route{
			Path: "/products",
			Method: http.MethodPost,
			Handler: r.productsController.PostProduct,
		},
		&Route{
			Path: "/products",
			Method: http.MethodGet,
			Handler: r.productsController.GetProducts,
		},
		&Route{
			Path: "/products/{product_id}",
			Method: http.MethodGet,
			Handler: r.productsController.GetProduct,
		},
		&Route{
			Path:    "/products/{product_id}",
			Method:  http.MethodPut,
			Handler: r.productsController.PutProduct,
		},
		&Route{
			Path:    "/products/{product_id}",
			Method:  http.MethodDelete,
			Handler: r.productsController.DeleteProduct,
		},
		&Route{
			Path:    "/search/products",
			Method:  http.MethodGet,
			Handler: r.productsController.SearchProducts,
		},
	}
}