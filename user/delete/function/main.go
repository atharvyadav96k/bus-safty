package org_register

import "net/http"

func UsersDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UsersDelete: User deletion was successful."))
}
