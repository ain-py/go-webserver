package auth

import (
	"errors"
	"net/http"
	"strings"
)

// * extract API key form headers of HTTP req
// * Authorization: ApiKey {api key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("not auth infor found")
	}
	vals := strings.Split(val, " ")
	//log.Default().Println(len(vals))
	if len(vals) != 2 {
		return "", errors.New("bad auth heaaders")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("bad first part of auth heaaders")
	}
	return vals[1], nil
}
