package controller

import (
	"encoding/json"
	"gopher-bank/error"
	"gopher-bank/model"
	"gopher-bank/service"
	"net/http"
)

var (
	userService service.UserService
)

type controller struct{}

type UserController interface {
	FindAllUsers(w http.ResponseWriter, r *http.Request)
	AddUser(w http.ResponseWriter, r *http.Request)
}

func NewUserController(service service.UserService) UserController {
	userService = service
	return &controller{}
}

func (c *controller) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.ServiceError{Message: "Error unmarshalling the request"})
	}

	err = userService.Validate(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.ServiceError{Message: err.Error()})
	}

	result, err := userService.Create(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.ServiceError{Message: "Error saving the post"})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (c *controller) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	posts, err := userService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.ServiceError{Message: "Error getting the posts"})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
