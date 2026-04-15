package org_register

import "net/http"

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GetAllUsers: Successfully retrieved all users."))
}
