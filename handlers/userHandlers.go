package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	users "www.github.com/Ammce/banking-go-auth/domain/Users"
)

type UserHandlers struct {
	service users.DefaultUserService
}

func (uh UserHandlers) Register(w http.ResponseWriter, r *http.Request) {
	var body users.User
	err := json.NewDecoder(r.Body).Decode(&body)
	fmt.Println("00")
	if err != nil {
		fmt.Println("11")
		writeReponse(w, http.StatusBadRequest, err)
	} else {
		fmt.Println("22")
		r, err := uh.service.Register(body)
		if err != nil {
			fmt.Println(*err)
			writeReponse(w, http.StatusBadRequest, err)
		} else {
			fmt.Println("44")
			writeReponse(w, http.StatusOK, r)
		}
	}
}

func NewUserHandlers(service users.DefaultUserService) UserHandlers {
	return UserHandlers{
		service: service,
	}
}
