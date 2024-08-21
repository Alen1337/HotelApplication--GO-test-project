package api_v1

import (
	"net/http"
)

func CheckAuthorization(w http.ResponseWriter, r *http.Request) (bool, string) {
	authKey := r.URL.Query().Get("auth")

	if authKey == "" {
		return false, ""
	}

	// TODO: Authorization check

	return true, authKey
}
