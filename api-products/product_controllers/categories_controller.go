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

type CategoriesController interface {
	PostCategory(http.ResponseWriter, *http.Request)
	GetCategory(http.ResponseWriter, *http.Request)
	GetCategories(http.ResponseWriter, *http.Request)
	PutCategory(http.ResponseWriter, *http.Request)
	DeleteCategory(http.ResponseWriter, *http.Request)
}

type categoriesControllerImpl struct {
	categoriesRepository product_repository.CategoriesRepository
}



func NewCategoriesController(categoriesRepository product_repository.CategoriesRepository) *categoriesControllerImpl {
	return &categoriesControllerImpl{categoriesRepository}
}

func (c *categoriesControllerImpl) PostCategory(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	category := &models.Category{}
	err = json.Unmarshal(bytes, category)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = category.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	category, err = c.categoriesRepository.Save(category)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}
	
	buildCreatedResponse(w, buildLocation(r, category.ID))
	utils.WriteAsJson(w, category)
}

func (c *categoriesControllerImpl) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.categoriesRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, categories)
}

func (c *categoriesControllerImpl) GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	category_id, err := strconv.ParseUint(params["category_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	category, err := c.categoriesRepository.Find(category_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteAsJson(w, category)
}

func (c *categoriesControllerImpl) PutCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	category_id, err := strconv.ParseUint(params["category_id"], 10, 64)
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

	category := &models.Category{}
	err = json.Unmarshal(bytes, category)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	category.ID = category_id

	err = category.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.categoriesRepository.Update(category)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *categoriesControllerImpl) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	category_id, err := strconv.ParseUint(params["category_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.categoriesRepository.Delete(category_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	
	buildDeleteResponse(w, category_id)
	utils.WriteAsJson(w, "{}")
}



