package auth

import (
	"app-profesor/backend/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Err struct {
	Err string `json:"error"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Err{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}

	user, err = a.Register(user)

	if err != nil {
		w.WriteHeader(http.StatusLocked)
		error := Err{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
	return

}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var login LoginRequest

	err := json.NewDecoder(r.Body).Decode(&login)

	log.Println(login)
	w.Header().Set("Content-Type", "application/json")

	user, err := a.Login(login.Username, login.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Err{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
