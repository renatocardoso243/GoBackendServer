package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey returns the API key from the request
//Example
//Authorization: insert <API_KEY>
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication header provided")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authentication header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of authentication header")
	}
	return vals[1], nil
}