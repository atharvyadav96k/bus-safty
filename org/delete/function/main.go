package org_register

import "net/http"

func OrgDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OrgDelete: Organization deletion was successful."))
}
