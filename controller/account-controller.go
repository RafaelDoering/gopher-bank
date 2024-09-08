package controller

import (
	"encoding/json"
	"gopher-bank/error"
	"gopher-bank/model"
	"gopher-bank/service"
	"net/http"
)

var (
	accountService service.AccountService
)

type accountController struct{}

type AccountController interface {
	FindAccountById(w http.ResponseWriter, r *http.Request)
	AddAccount(w http.ResponseWriter, r *http.Request)
}

func NewAccountController(service service.AccountService) AccountController {
	accountService = service
	return &accountController{}
}

func (c *accountController) AddAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var account model.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.ServiceError{Message: "Error unmarshalling the request"})
	}

	err = accountService.Validate(&account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.ServiceError{Message: err.Error()})
	}

	result, err := accountService.Create(&account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.ServiceError{Message: "Error saving the account"})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (c *accountController) FindAccountById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	account, err := accountService.FindById(1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.ServiceError{Message: "Error getting the account"})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}
