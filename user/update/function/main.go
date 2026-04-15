package org_register

import "net/http"

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UpdateUser: User update was successful."))
}
