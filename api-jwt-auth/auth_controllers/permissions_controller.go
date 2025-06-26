package auth_controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/auth_repository"
	"github.com/austineyoyogie/go-hardware-store/api-jwt-auth/models"
	"github.com/austineyoyogie/go-hardware-store/utils"
	"github.com/gorilla/mux"
)

// https://www.youtube.com/watch?v=9tp82wyG-Mc&list=PLkVx132FdJZlTc_1gucKZ00b_s45DQlVQ&index=5

type PermissionsController interface {
	PostPermission(http.ResponseWriter, *http.Request)
	GetPermission(http.ResponseWriter, *http.Request)
	GetPermissions(http.ResponseWriter, *http.Request)
	PutPermission(http.ResponseWriter, *http.Request)
	DeletePermission(http.ResponseWriter, *http.Request)
}

type permissionsControllerImpl struct {
	permissionsRepository auth_repository.PermissionsRepository
}

func NewPermissionsController(permissionsRepository auth_repository.PermissionsRepository) *permissionsControllerImpl {
	return &permissionsControllerImpl{permissionsRepository}
}

func (p *permissionsControllerImpl) PostPermission(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
	}
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	permission := &models.Permission{}
	err = json.Unmarshal(bytes, permission)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = permission.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	permission, err = p.permissionsRepository.Save(permission)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, permission.ID))
	utils.WriteAsJson(w, permission)
}

func (p *permissionsControllerImpl) GetPermission(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	permission_id, err := strconv.ParseUint(params["permission_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	permission, err := p.permissionsRepository.Find(permission_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, permission)
}

func (p *permissionsControllerImpl) GetPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := p.permissionsRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, permissions)
}

func (p *permissionsControllerImpl) PutPermission(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	permission_id, err := strconv.ParseUint(params["permission_id"], 10, 64)
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

	permission := &models.Permission{}
	err = json.Unmarshal(bytes, permission)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	permission.ID = permission_id

	err = permission.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = p.permissionsRepository.Update(permission)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (p *permissionsControllerImpl) DeletePermission(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	permission_id, err := strconv.ParseUint(params["permission_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = p.permissionsRepository.Delete(permission_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	buildDeleteResponse(w, permission_id)
	utils.WriteAsJson(w, "{}")
}
