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

type UsersController interface {
	PostUser(http.ResponseWriter, *http.Request)
	VerifyUser(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)	
	GetUsers(http.ResponseWriter, *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	ResetPasswordUser(w http.ResponseWriter, r *http.Request)
	GetNewPasswordUser(w http.ResponseWriter, r *http.Request)
	PutNewPasswordUser(w http.ResponseWriter, r *http.Request)
}

type usersControllerImpl struct {
	usersRepository auth_repository.UsersRepository
}

func NewUsersController(usersRepository auth_repository.UsersRepository) *usersControllerImpl {
	return &usersControllerImpl{usersRepository}
}

func (u *usersControllerImpl) PostUser(w http.ResponseWriter, r *http.Request) {
	
	if r.Body != nil { 
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user := &models.User{}
	err = json.Unmarshal(bytes, user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = user.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	
	user, err = u.usersRepository.Save(user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, user.ID))
	utils.WriteAsJson(w, user)
}

func (u *usersControllerImpl) VerifyUser(w http.ResponseWriter, r *http.Request) {
	
	email := r.URL.Query().Get("email")
	token := r.URL.Query().Get("token")

	_, err := u.usersRepository.Verify(email, token)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteAsJson(w, nil)
}

func (u *usersControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
	user_id, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	user, err := u.usersRepository.Find(user_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteAsJson(w, user)
}

func (u *usersControllerImpl) GetUsers(w http.ResponseWriter, r *http.Request) {
	
 	users, err := u.usersRepository.FindAll()
 	if err != nil {
 		utils.WriteError(w, err, http.StatusInternalServerError)
 		return
 	}
 	utils.WriteAsJson(w, users)
}

func (u *usersControllerImpl) PutUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user_id, err := strconv.ParseUint(params["user_id"], 10, 64)
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

	user := &models.User{}
	err = json.Unmarshal(bytes, user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user.ID = user_id

	err = user.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = u.usersRepository.Update(user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (u *usersControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
	user_id, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}
	
	err = u.usersRepository.Delete(user_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	
	buildDeleteResponse(w, user_id)
	utils.WriteAsJson(w, "{}")
}

// PUT localhost:8000/reset
func (u *usersControllerImpl) ResetPasswordUser(w http.ResponseWriter, r *http.Request) {
	
	if r.Body != nil { 
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user := &models.User{}
	err = json.Unmarshal(bytes, user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}
	 // create a single email validation

	// err = user.Validate()
	//   if err != nil {
	//   	utils.WriteError(w, err, http.StatusBadRequest)
	//   	return
	// }

	// Check if user email exists
	user, err = u.usersRepository.FindEmail(user.Email)
	if err != nil {
	utils.WriteError(w, err, http.StatusUnprocessableEntity)
	 	return 
	} 

	err = u.usersRepository.ResetToken(user)
	if err != nil {
	utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}
	
	buildCreatedResponse(w, buildLocation(r, user.ID))
	utils.WriteAsJson(w, user)
}

func (u *usersControllerImpl) GetNewPasswordUser(w http.ResponseWriter, r *http.Request) {
	
	email := r.URL.Query().Get("email")
	token := r.URL.Query().Get("token")

	user, err := u.usersRepository.FindPasswordResetUser(email, token)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteAsJson(w, user)
}

func (u *usersControllerImpl) PutNewPasswordUser(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil { 
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user := &models.User{}
	err = json.Unmarshal(bytes, user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = user.PutNewPasswordUserValidate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
	 	return
	}

	// check if user email does match
	user, err = u.usersRepository.FindPasswordResetUser(user.Email, user.Token)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	
	err = u.usersRepository.UpdateNewPasswordUser(user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.WriteAsJson(w, map[string]bool{"success": err == nil})	
}

