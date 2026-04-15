package org_register

import "net/http"

func GetAllOrg(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GetAllOrg: Successfully retrieved all organizations."))
}
