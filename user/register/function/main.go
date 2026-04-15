package user_register

import "net/http"

func UserRegister(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UserRegister: User registration was successful."))
}
