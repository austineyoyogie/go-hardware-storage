package product_controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/austineyoyogie/go-hardware-store/api-products/models"
	"github.com/austineyoyogie/go-hardware-store/api-products/product_repository"
	"github.com/austineyoyogie/go-hardware-store/utils"
	"github.com/gorilla/mux"
)

type ProductsController interface {
	PostProduct(http.ResponseWriter, *http.Request);
	GetProduct(http.ResponseWriter, *http.Request)
	GetProducts(http.ResponseWriter, *http.Request)
	PutProduct(http.ResponseWriter, *http.Request)
	DeleteProduct(http.ResponseWriter, *http.Request)
	SearchProducts(http.ResponseWriter, *http.Request)
}

type productsControllerImpl struct {
	productsRepository product_repository.ProductsRepository
	paginationBuilder  product_repository.PaginationBuilderRepository
}

func NewProductsController(productsRepository product_repository.ProductsRepository) *productsControllerImpl {

	return &productsControllerImpl{productsRepository,
		product_repository.NewPaginationBuilderRepository(productsRepository)}
}

func (c *productsControllerImpl) PostProduct(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	product := &models.Product{}
	err = json.Unmarshal(bytes, product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = product.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	product, err = c.productsRepository.Save(product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, product.ID))
	utils.WriteAsJson(w, product)
}

func (c *productsControllerImpl) GetProducts(w http.ResponseWriter, r *http.Request) {

	meta, err := c.paginationBuilder.BuildProductsMetadata(r)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	/*
	// if you don't need pagination use this
	products, err := c.productsRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.WriteAsJson(w, products)
	*/

	// http://localhost:5000/products?page=5

	// http://localhost:5000/products?page=3&limit=3

	// http://localhost:5000/products?page=9&limit=7

	elems, err := c.productsRepository.Paginate(meta)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, elems)
}

func (c *productsControllerImpl) GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	product_id, err := strconv.ParseUint(params["product_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	product, err := c.productsRepository.Find(product_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteAsJson(w, product)
}


func (c *productsControllerImpl) PutProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	product_id, err := strconv.ParseUint(params["product_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	product := &models.Product{}
	err = json.Unmarshal(bytes, product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	product.ID = product_id

	err = product.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	//product.CheckStatus()

	err = c.productsRepository.Update(product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *productsControllerImpl) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	product_id, err := strconv.ParseUint(params["product_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.productsRepository.Delete(product_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	buildDeleteResponse(w, product_id)
	utils.WriteAsJson(w, "{}")
}

func (c *productsControllerImpl) SearchProducts(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	q := params.Get("q")

	if q == "" {
		utils.WriteAsJson(w, "[]")
		return
	}

	// http://localhost:5000/search/products?q=GTX
	
	products, err := c.productsRepository.Search(q)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteAsJson(w, products)
}
