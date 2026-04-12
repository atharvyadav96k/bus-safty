package request

import (
	"encoding/json"
	"net/http"
)

func Decoder[T any](r *http.Request) (T, error) {
	var body T
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return body, err
	}
	defer r.Body.Close()
	return body, nil
}
