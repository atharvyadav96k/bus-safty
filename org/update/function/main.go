package org_register

import "net/http"

func UpdateOrg(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UpdateOrg: Organization update was successful."))
}
