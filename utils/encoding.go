package utils

import (
	"encoding/json"
	"net/http"

	"github.com/amieldelatorre/myip/customerrors"
)

func EncodeResponse[T any](w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return &customerrors.EncodeResponseError{}
	}

	return nil
}
