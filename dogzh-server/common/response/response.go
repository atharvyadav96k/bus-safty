package response

import (
	"encoding/json"
	"net/http"
)

func (r *Response) SendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	json.NewEncoder(w).Encode(r)
}
