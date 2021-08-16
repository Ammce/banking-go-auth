package handlers

import (
	"encoding/json"
	"net/http"

	users "www.github.com/Ammce/banking-go-auth/domain/Users"
)

type UserHandlers struct {
	service users.DefaultUserService
}

func (uh UserHandlers) Register(w http.ResponseWriter, r *http.Request) {
	var body users.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeReponse(w, http.StatusBadRequest, err)
	} else {
		r, err := uh.service.Register(body)
		if err != nil {
			writeReponse(w, http.StatusBadRequest, err)
		} else {
			writeReponse(w, http.StatusOK, r)
		}
	}
}

func NewUserHandlers(service users.DefaultUserService) UserHandlers {
	return UserHandlers{
		service: service,
	}
}
