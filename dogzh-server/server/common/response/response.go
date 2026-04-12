package response

import (
	"encoding/json"
	"net/http"
)

func (r *Response) sendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	json.NewEncoder(w).Encode(r)
}

func HttpResponseCreated(w http.ResponseWriter, message string, Data any) {
	res := Response{
		Message: message,
		Status:  http.StatusCreated,
		Data:    Data,
	}
	res.sendResponse(w)
}

func HttpResponseOK(w http.ResponseWriter, message string, Data any) {
	res := Response{
		Message: message,
		Status:  http.StatusOK,
		Data:    Data,
	}
	res.sendResponse(w)
}

func HttpResponseBadRequest(w http.ResponseWriter, message string) {
	res := Response{
		Message: message,
		Status:  http.StatusBadRequest,
	}
	res.sendResponse(w)
}

func HttpResponseForbidden(w http.ResponseWriter, message string) {
	res := Response{
		Status:  http.StatusForbidden,
		Message: "You don't have access to the resource",
	}
	res.sendResponse(w)
}

func HttpResponseUnauthorized(w http.ResponseWriter) {
	res := Response{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access",
	}
	res.sendResponse(w)
}
