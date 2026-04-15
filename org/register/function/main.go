package org_register

import "net/http"

func OrgRegister(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Organization registration successful"))
}
